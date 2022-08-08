package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	ID      string  `json:"ID"`
	Name    string  `json:"Name"`
	TaxRate float64 `json:"Tax Rate"`
	Price   float64 `json:"Price"`
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	// Check user logging or not
	if ActiveUser == &emptyUser || ActiveUser == nil {
		fmt.Fprintf(w, "There is no active user. Please assign or login user!")
		return
	}

	// Take product from SQL database
	singleProduct, err := GetProductFromDatabase(productId)
	if err != nil {
		errMessage := fmt.Sprint(err)
		fmt.Fprintf(w, errMessage)
		return
	}

	if len(Basket) == 0 {
		// Add product if basket is empty
		AddedProduct := ProductInBasket{singleProduct, 1, orderID}
		Basket = append(Basket, AddedProduct)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(singleProduct)
		return
	} else {
		singleProductInBasket := GetProductFromBasket(singleProduct)
		if singleProductInBasket != EmptyProductInBasket {
			// If product is also available in basket
			IncreaseAmount(singleProductInBasket)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Basket)
			return
		} else {
			// If product is not available in basket
			AddedProduct := ProductInBasket{singleProduct, 1, orderID}
			Basket = append(Basket, AddedProduct)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Basket)
			return
		}
	}
}

func GetProductFromDatabase(productId string) (Product, error) {
	db, err := ConnectSQL("ordersdb")

	if err != nil {
		log.Fatal(err)
		return Product{}, err
	}
	res, err := db.Query("SELECT * FROM products WHERE id = ?", productId)

	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {

		var product Product
		err := res.Scan(&product.ID, &product.Name, &product.TaxRate, &product.Price)

		if err != nil {
			log.Fatal(err)
		}

		return product, nil
	}
	return Product{}, errors.New(fmt.Sprintf("a product with this ID %v does not exist in the product database", productId))
}

func GetProductFromBasket(singleProduct Product) ProductInBasket {
	for _, singleProductInBasket := range Basket {
		if singleProductInBasket.Product == singleProduct {
			return singleProductInBasket
		}
	}
	return ProductInBasket{}
}

func IncreaseAmount(productInBasket ProductInBasket) {
	productInBasketId := productInBasket.Product.ID
	//db, _ := ConnectSQL("ordersdb")
	for index, singleProductInBasket := range Basket {
		if singleProductInBasket.Product.ID == productInBasketId {
			singleProductInBasket.Amount = singleProductInBasket.Amount + 1
			Basket = append(Basket[:index], singleProductInBasket)
		}
	}
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

func dropProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var updatedProductInBasket ProductInBasket
	for index, singleProductInBasket := range Basket {
		if singleProductInBasket.Product.ID == productId {
			updatedProductInBasket.Product = singleProductInBasket.Product
			updatedProductInBasket.Amount = singleProductInBasket.Amount - 1
			if updatedProductInBasket.Amount == 0 {
				deleteProduct(w, r)
				return
			}
			Basket[index] = updatedProductInBasket
			json.NewEncoder(w).Encode(updatedProductInBasket)
			return
		}
	}
	fmt.Fprintf(w, "There is no product with this ID %v in the basket", productId)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if len(Basket) == 0 {
		fmt.Fprintf(w, "There is no product in basket!")
		return
	} else {
		for index, singleProductInBasket := range Basket {
			if singleProductInBasket.Product.ID == productId {
				Basket = append(Basket[:index], Basket[index+1:]...)
				fmt.Fprintf(w, "The product with ID %v has been deleted successfully", productId)
				return
			}
		}
		fmt.Fprintf(w, "There is no product with this ID %v in the basket", productId)
	}
}
