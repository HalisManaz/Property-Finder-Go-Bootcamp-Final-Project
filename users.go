package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ActiveUser *User

type User struct {
	ID           string  `json:"ID"`
	UserName     string  `json:"UserName"`
	Type         string  `json:"Type"`
	Streak       int     `json:"Streak"`
	MonthlyTotal float64 `json:"MonthlyTotal"`
}

func (u *User) PaymentDone(Basket basketProducts, TotalPrice, TaxAmount, Discount, AmountDue float64) {
	db, _ := ConnectSQL("ordersdb")
	for _, productInBasket := range Basket {
		_, err := db.Query("INSERT INTO basket(ID, Name,TaxRate,Price,Amount,OrderID) VALUES (?,?,?,?,?,?)",
			productInBasket.Product.ID, productInBasket.Product.Name, productInBasket.Product.TaxRate,
			productInBasket.Product.Price, productInBasket.Amount, productInBasket.OrderID)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err := db.Query("INSERT INTO receipts(TotalPrice,TaxAmount,Discount,AmountDue,OrderID) VALUES (?,?,?,?,?)",
		TotalPrice, TaxAmount, Discount, AmountDue, orderID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("INSERT INTO userandorder(ID,orderID) VALUES (?,?)", ActiveUser.ID, orderID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("UPDATE users SET Type = ?,Streak = ? ,MonthlyTotal = ? WHERE ID = ?;",
		ActiveUser.Type, ActiveUser.Streak, ActiveUser.MonthlyTotal, ActiveUser.ID)
	if err != nil {
		log.Fatal(err)
	}
}
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid input for creating user")
	}
	json.Unmarshal(reqBody, &newUser)
	newUser = newUser.CreateUser(newUser.UserName)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	db, err := ConnectSQL("ordersdb")
	_, err = db.Query("INSERT INTO users(ID, UserName,Type,Streak,MonthlyTotal) VALUES (?,?,?,?,?)", newUser.ID, newUser.UserName, newUser.Type, newUser.Streak, newUser.MonthlyTotal)

	if err != nil {
		panic(err.Error())
	}
}

func ListAllUser(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectSQL("ordersdb")
	res, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var user User
		err := res.Scan(&user.ID, &user.UserName, &user.Type, &user.Streak, &user.MonthlyTotal)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(user)
	}

}
