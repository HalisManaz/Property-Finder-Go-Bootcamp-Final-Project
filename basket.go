package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ProductInBasket struct {
	Product Product `json:"Product"`
	Amount  int     `json:"Amount"`
	OrderID int     `json:"OrderID"`
}

type UserAndOrder struct {
	UserID  string `json:"UserID"`
	OrderID int    `json:"OrderID"`
}

type basketProducts []ProductInBasket

var EmptyProductInBasket ProductInBasket

var Basket = basketProducts{}

func ListAllProductsInBasket(w http.ResponseWriter, r *http.Request) {
	var TotalPrice, TaxAmount, Discount, AmountDue float64
	if ActiveUser == &emptyUser || ActiveUser == nil {
		fmt.Fprintf(w, "There is no active user. Please assign or login user!")
		return
	}
	//ActiveUser.CheckUserDiscount()
	Discount = CalculateDiscount(Basket, ActiveUser, "Check")
	fmt.Fprintf(w, "BASKET\n---------------------------------------\n")
	json.NewEncoder(w).Encode(Basket)
	for _, item := range Basket {
		TotalPrice += item.Product.Price * float64(item.Amount)
		TaxAmount += item.Product.Price * float64(item.Amount) * (item.Product.TaxRate / 100)
	}
	fmt.Fprintf(w, "RECEIPT\n---------------------------------------\n")
	AmountDue = TotalPrice - Discount
	fmt.Fprintf(w, "Total Amount: %.2f\nDiscount: %.2f\nTax Amount: %.2f\nAmount Due: %.2f",
		TotalPrice, Discount, TaxAmount, AmountDue)
}

func AllPastOrders(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	if ActiveUser == &emptyUser || ActiveUser == nil {
		fmt.Fprintf(w, "There is no active user. Please assign or login user!")
		return
	}

	db, err := ConnectSQL("ordersdb")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := db.Query("SELECT * FROM userandorder WHERE ID = ?", userID)

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var userAndOrder UserAndOrder
		var pastReceipt Receipts
		var pastBasketProduct ProductInBasket
		var pastBasket basketProducts

		if err := res.Scan(&userAndOrder.UserID, &userAndOrder.OrderID); err != nil {
			log.Fatal(err)
		}

		receipt, err := db.Query("SELECT * FROM receipts WHERE OrderID = ?", userAndOrder.OrderID)
		basket, err := db.Query("SELECT * FROM basket WHERE OrderID = ?", userAndOrder.OrderID)

		for receipt.Next() {
			for basket.Next() {
				if err := basket.Scan(&pastBasketProduct.Product.ID, &pastBasketProduct.Product.Name, &pastBasketProduct.Product.TaxRate,
					&pastBasketProduct.Product.Price, &pastBasketProduct.Amount, &pastBasketProduct.OrderID); err != nil {
					log.Fatal(err)
				}
				pastBasket = append(pastBasket, pastBasketProduct)
			}
			if err := receipt.Scan(&pastReceipt.TotalPrice, &pastReceipt.TaxAmount, &pastReceipt.Discount,
				&pastReceipt.AmountDue, &pastReceipt.OrderID); err != nil {
				log.Fatal(err)
			}

		}
		json.NewEncoder(w).Encode(Order{pastBasket, pastReceipt})

		if err != nil {
			log.Fatal(err)
		}
	}

	if err = res.Err(); err == nil {
		return
	}

	fmt.Fprintf(w, "There is no record of this user!\n")
}
