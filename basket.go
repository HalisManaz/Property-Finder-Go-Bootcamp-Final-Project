package main

type ProductInBasket struct {
	Product Product `json:"Product"`
	Amount  int     `json:"Amount"`
	OrderID int     `json:"OrderID"`
}

type basketProducts []ProductInBasket

var EmptyProductInBasket ProductInBasket

var Basket = basketProducts{}
