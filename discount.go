package main

import (
	"golang.org/x/exp/slices"
	"math"
	"sort"
)

func calculateDiscount(Basket basketProducts, u *User, paymentOrCheck string) (discountAmount float64) {
	var allDiscounts = []float64{0.0}
	var discount, basketTotalPrice float64

	// Total price of basket
	for _, item := range Basket {
		basketTotalPrice += item.Product.Price * float64(item.Amount)
	}

	streakAmount := 100.0
	streakDiscountsProductIDs := []string{"1", "4", "7"}

	// Part a
	if (u.Streak+1)%4.0 == 0 && u.Streak != 0 {
		for _, productInBasket := range Basket {
			if slices.Contains(streakDiscountsProductIDs, productInBasket.Product.ID) {
				if productInBasket.Product.TaxRate == 18 {
					streakDiscount := productInBasket.Product.Price * 0.15 * float64(productInBasket.Amount)
					allDiscounts = append(allDiscounts, math.Round(streakDiscount*100)/100)
				} else if productInBasket.Product.TaxRate == 8 {
					streakDiscount := productInBasket.Product.Price * 0.10 * float64(productInBasket.Amount)
					allDiscounts = append(allDiscounts, math.Round(streakDiscount*100)/100)
				}
			}
		}
	}

	if paymentOrCheck == "Payment" {
		if basketTotalPrice > streakAmount {
			(*u).Streak++
		}
	}

	// Part b
	for _, productInBasket := range Basket {
		if productInBasket.Amount > 3 {
			discount = (float64(productInBasket.Amount) - 3) * productInBasket.Product.Price * 0.08
			allDiscounts = append(allDiscounts, discount)
		}
	}

	// Part c
	mountlyLimit := 500.0
	if u.MonthlyTotal > mountlyLimit {
		(*u).Type = "Premium"
	}

	if paymentOrCheck == "Payment" {
		(*u).MonthlyTotal += basketTotalPrice
	}

	if u.Type == "Premium" {
		discount = basketTotalPrice * 0.1
		allDiscounts = append(allDiscounts, discount)
	}
	// Part d OK!
	sort.Float64s(allDiscounts)
	return allDiscounts[len(allDiscounts)-1]
}
