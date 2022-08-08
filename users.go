package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	ID           string  `json:"ID"`
	UserName     string  `json:"UserName"`
	Type         string  `json:"Type"`
	Streak       int     `json:"Streak"`
	MonthlyTotal float64 `json:"MonthlyTotal"`
}

type Receipts struct {
	TotalPrice float64 `json:"TotalPrice"`
	TaxAmount  float64 `json:"TaxAmount"`
	Discount   float64 `json:"Discount"`
	AmountDue  float64 `json:"AmountDue"`
	OrderID    int     `json:OrderID`
}

type Order struct {
	Basket  basketProducts `json:"Basket"`
	Receipt Receipts       `json:"Receipt"`
}

type PastOrder struct {
	Order []Order `json:"Past Orders"`
}

var ActiveUser *User

func (u User) CreateUser(UserName string) (user User) {
	rand.Seed(time.Now().UnixNano())
	u.ID = fmt.Sprint(rand.Intn(1000))
	u.UserName = UserName
	u.Type = "Normal"
	u.Streak = 0
	u.MonthlyTotal = 0.0
	return u
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

func SetActiveUser(w http.ResponseWriter, r *http.Request) {
	ActiveUserID := mux.Vars(r)["id"]
	db, err := ConnectSQL("ordersdb")

	res, err := db.Query("SELECT * FROM users WHERE id = ?", ActiveUserID)
	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {

		var user User
		err := res.Scan(&user.ID, &user.UserName, &user.Type, &user.Streak, &user.MonthlyTotal)

		if err != nil {
			log.Fatal(err)
		}

		if ActiveUser != nil {
			if ActiveUser.ID == user.ID {
				fmt.Fprintf(w, "Active user already selected user!")
				return
			} else {
				orderID = rand.Intn(10000)
				Basket = basketProducts{}
			}
		}

		ActiveUser = &user
		rand.Seed(time.Now().UnixNano())
		orderID = rand.Intn(100_000_000)
		json.NewEncoder(w).Encode(ActiveUser)
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
