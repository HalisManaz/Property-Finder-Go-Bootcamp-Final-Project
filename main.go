package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var orderID int
var emptyUser User

// Homepage
func homeLink(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome PF Market\n"+
		"If there are more than 3 items of the same product, then fourth and subsequent ones would have 8 percent off.\n"+
		"If you made purchase which is more than 500 in a month then all subsequent purchaseshave 10 percen off.\n"+
		"Every fourth order whose total is more than 100 may have discount depending on products which are discount products of the month. "+
		"Products whose VAT is 1 percent don’t have any discount but products whose VAT is 8 percent and 18 percent have discount of 10 percent and 15 percent respectively.\n"+
		"Only one discount can be applied at a time and only disconunt which is the highest are applied")
	if err != nil {
		return
	}
}

func Payment(w http.ResponseWriter, r *http.Request) {
	var TotalPrice, TaxAmount, Discount, AmountDue float64
	paymentDecide := mux.Vars(r)["decide"]
	// When payment declined
	if strings.ToLower(paymentDecide) == "y" {
		if len(Basket) == 0 {
			_, _ = fmt.Fprint(w, "Basket is empty! Please fill the basket to pay")
			return
		} else {
			// When payment accepted
			// Check user is logging or not
			if ActiveUser == &emptyUser || ActiveUser == nil {
				_, _ = fmt.Fprint(w, "There is no active user. Please assign or login user!")
				return
			} else {
				// If user is logging and payment accepted, payment are done.
				_, _ = fmt.Fprint(w, "Payment are done!\n")
				Discount = calculateDiscount(Basket, ActiveUser, "Payment")
				_, _ = fmt.Fprintf(w, "BASKET\n---------------------------------------\n")
				err := json.NewEncoder(w).Encode(Basket)

				if err != nil {
					return
				}

				for _, item := range Basket {
					TotalPrice += item.Product.Price * float64(item.Amount)
					TaxAmount += item.Product.Price * float64(item.Amount) * (item.Product.TaxRate / 100)
				}
				_, _ = fmt.Fprintf(w, "RECEIPT\n---------------------------------------\n")
				AmountDue = TotalPrice - Discount
				_, _ = fmt.Fprintf(w, "Total Amount: %.2f\nDiscount: %.2f\nTax Amount: %.2f\nAmount Due: %.2f", TotalPrice, Discount, TaxAmount, AmountDue)
				ActiveUser.PaymentDone(Basket, TotalPrice, TaxAmount, Discount, AmountDue)
				Basket = basketProducts{}
				rand.Seed(time.Now().UnixNano())
				orderID = rand.Intn(100_000_000)
				return
			}
		}
	} else if strings.ToLower(paymentDecide) == "n" {
		_, err := fmt.Fprint(w, "Payment declined")
		if err != nil {
			return
		}
		return
	}
}

func ConnectSQL(sqldb string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:sqlpassword@tcp(127.0.0.1:3306)/"+sqldb)
	//defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/products", ListAllProducts).Methods("GET")
	router.HandleFunc("/product/{id}", addProduct).Methods("POST")
	router.HandleFunc("/basket/{id}", dropProduct).Methods("PATCH")
	router.HandleFunc("/basket/{id}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/basket", ListAllProductsInBasket).Methods("GET")
	router.HandleFunc("/payment/{decide}", Payment).Methods("GET")
	router.HandleFunc("/users", ListAllUser).Methods("GET")
	router.HandleFunc("/user", CreateNewUser).Methods("POST")
	router.HandleFunc("/setActiveUser/{id}", SetActiveUser).Methods("GET")
	router.HandleFunc("/pastOrders/{id}", AllPastOrders).Methods("GET")

	log.Fatalln(http.ListenAndServe(":8080", router))
}
