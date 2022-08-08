package main

import "testing"

var testVariables = []struct {
	Basket         basketProducts
	User           *User
	PaymentOrCheck string
	Discount       float64
}{
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 3, 1}},
		&User{"1", "TestUserA", "Normal", 0, 0}, "Payment", 0},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 8, 1}},
		&User{"1", "TestUserA", "Normal", 0, 0}, "Payment", 2.0},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 3, 1}},
		&User{"2", "TestUserB", "Premium", 0, 700}, "Payment", 1.5},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 8, 1}},
		&User{"2", "TestUserB", "Premium", 0, 700}, "Payment", 4.0},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 3, 1}},
		&User{"3", "TestUserC", "Normal", 3, 0}, "Payment", 2.25},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 8, 1}},
		&User{"3", "TestUserC", "Normal", 3, 0}, "Payment", 6.0},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 3, 1}},
		&User{"4", "TestUserD", "Premium", 3, 700}, "Payment", 2.25},
	{basketProducts{ProductInBasket{Product{"1", "Chocolate", 18, 5.0}, 8, 1}},
		&User{"4", "TestUserD", "Premium", 3, 700}, "Payment", 6.0},
}

/* Note for reader:
TestUserA -> No premium customer and no streak available
TestUserB -> Premium customer and no streak available
TestUserC -> No premium customer and streak available
TestUserD -> premium customer and order streak available

Also, there are two type of basket:
A -> 3 x Chocolate
Basket B -> 8 x Chocolate

Both two type of basket are test for all users and function works properly.
*/

func TestCalculateDiscount(t *testing.T) {
	for _, testVariable := range testVariables {
		got := CalculateDiscount(testVariable.Basket, testVariable.User, testVariable.PaymentOrCheck)
		want := testVariable.Discount
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	}
}
