package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductInBasket struct {
	Product Product `json:"Product"`
	Amount  int     `json:"Amount"`
	OrderID int     `json:"OrderID"`
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
	Discount = calculateDiscount(Basket, ActiveUser, "Check")
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
