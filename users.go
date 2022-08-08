package main

type User struct {
	ID           string  `json:"ID"`
	UserName     string  `json:"UserName"`
	Type         string  `json:"Type"`
	Streak       int     `json:"Streak"`
	MonthlyTotal float64 `json:"MonthlyTotal"`
}
