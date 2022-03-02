package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
)

func Routing() *echo.Echo {
	// user controller
	userController := controllers.NewUserController(NewSqlHandler())
	// echo instance
    e := echo.New()
    // routing
    // e.POST("/login",)
    // g := e.Group("/admin")
	// g.Use(FirebaseAuth)
	
	e.POST("/signup", userController.SignUp)
	// e.POST("/signin", userController.SignUp)

	e.GET("/users/:id", userController.GetUser)
    e.GET("/users",userController.GetAllUsers)
	return e
}
