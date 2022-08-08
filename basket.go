package main

type ProductInBasket struct {
	Product Product `json:"Product"`
	Amount  int     `json:"Amount"`
	OrderID int     `json:"OrderID"`
}
