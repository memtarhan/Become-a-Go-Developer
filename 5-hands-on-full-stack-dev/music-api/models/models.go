package models

import "time"

type Product struct {
	Image       string  `json:"image"`
	ImageAlt    string  `json:"imageAlt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
}

type Customer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	LoggedIn  bool   `json:"loggedIn"`
}

type Order struct {
	Product
	Customer
	CustomerId   int       `json:"customerId"`
	ProductId    int       `json:"productId"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchaseDate"`
}
