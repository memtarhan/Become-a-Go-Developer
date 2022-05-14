package rest

import (
	"github.com/gin-gonic/gin"
	"music-api/db"
	"music-api/models"
	"net/http"
	"strconv"
)

type HandlerInterface interface {
	GetProducts(context *gin.Context)
	GetPromos(context *gin.Context)
	AddUser(context *gin.Context)
	SignIn(context *gin.Context)
	SignOut(context *gin.Context)
	GetOrders(context *gin.Context)
	Charge(context *gin.Context)
}

type Handler struct {
	db db.DBLayer
}

func NewHandler() (*Handler, error) {
	// This creates a new pointer to the Handler object
	return new(Handler), nil
}

func (handler *Handler) GetProducts(context *gin.Context) {
	if handler.db == nil {
		return
	}

	products, err := handler.db.GetAllProducts()

	if err != nil {
		// First argument is the http status code, whereas the second argument is the body of the request
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//product := models.Product{}
	//product.ProductName = "Product name is this"
	//newProducts := []models.Product{product}

	context.JSON(http.StatusOK, products)
}

func (handler *Handler) GetPromos(context *gin.Context) {
	if handler.db == nil {
		return
	}

	promos, err := handler.db.GetPromos()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, promos)
}

func (handler *Handler) SignIn(context *gin.Context) {
	if handler.db == nil {
		return
	}

	var customer models.Customer
	err := context.ShouldBindJSON(&customer)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: - Check this out
	customer, err = handler.db.SignInUser(customer.Email, "")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, customer)
}

func (handler *Handler) AddUser(context *gin.Context) {
	if handler.db == nil {
		return
	}

	var customer models.Customer
	err := context.ShouldBindJSON(&customer)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err = handler.db.AddUser(customer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, customer)
}

func (handler *Handler) SignOut(context *gin.Context) {
	if handler.db == nil {
		return
	}

	p := context.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.db.SignOutUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (handler *Handler) GetOrders(context *gin.Context) {
	if handler.db == nil {
		return
	}

	// get id parameter
	p := context.Param("id")
	// convert the string 'p' to integer 'id'
	id, err := strconv.Atoi(p)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call the db layer method to get orders from id
	orders, err := handler.db.GetCustomerOrdersById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, orders)
}

func (handler *Handler) Charge(context *gin.Context) {
	if handler.db == nil {
		return
	}
}
