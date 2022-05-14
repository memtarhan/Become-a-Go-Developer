package db

/*
	Get a list of all products
	Get a list of all promotions
	Get a customer by the customer's first and last name
	Get a customer by the customer's id
	Get a product by the product's id
	Add a user to the database
	Mark a user in the database as signed in
	Mark a user in the database as signed out
	Get a list of customer orders by the customer's id
*/

import "music-api/models"

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerById(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(customer models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersById(int) ([]models.Order, error)
}
