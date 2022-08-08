package main

import "log"

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
