package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {

	// Define a handler
	handler, err := NewHandler()
	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, handler)
}

func RunAPIWithHandler(address string, handler HandlerInterface) error {
	// Get gin's default engine
	r := gin.Default()

	// get products
	r.GET("/products", handler.GetProducts)

	// get promos
	r.GET("/promos", handler.GetPromos)

	/*
		// post user sign in
		r.POST("/user/signIn", handler.SignIn)

		// post add user
		r.POST("/user", handler.AddUser)

		// post user sign out
		/*
				In the path below, our relative url needs to include the user id
			  	Since the id will differ based on the user, the Gin framework allows us to
			  	include a wildcard. In Gin, the wildcard will take the form ':id' to indicate
			  	that we are expecting a parameter here with the name 'id'
		*
		r.POST("/user/:id/signOut", handler.SignOut)

		// get user orders
		r.GET("/user/:id/orders", handler.GetOrders)

		// post purchase charge
		r.POST("/user/charge", handler.Charge)

	*/

	// Since a number of our routes start with /user/ and /users,
	// the preceding code can be refactored further using a method called Group():
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signOut", handler.SignOut)
		userGroup.GET("/:id/orders", handler.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", handler.Charge)
		usersGroup.POST("/signIn", handler.SignIn)
		usersGroup.POST("", handler.AddUser)
	}

	return r.Run(address)

	/*
		Observe the last line in our function: r.Run(address). We must call this method
		after we finish defining our API routes and handlers, so that our RESTful API
		server starts listening to incoming requests from HTTP clients.
	*/
}
