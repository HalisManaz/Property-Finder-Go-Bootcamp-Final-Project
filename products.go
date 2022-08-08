package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID      string  `json:"ID"`
	Name    string  `json:"Name"`
	TaxRate float64 `json:"Tax Rate"`
	Price   float64 `json:"Price"`
}

func ListAllProducts(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectSQL("ordersdb")
	rows, err := db.Query("SELECT * FROM products;")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer rows.Close()

	var productDatabase []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.TaxRate, &product.Price)
		if err != nil {
			fmt.Println(err)
			return
		}

		productDatabase = append(productDatabase, product)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(productDatabase); err != nil {
		fmt.Println(err)
	}
	return
}
