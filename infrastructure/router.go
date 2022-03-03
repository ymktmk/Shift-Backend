package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
    e := echo.New()
    // e.POST("/login",)
    // g := e.Group("/admin")
	// g.Use(FirebaseAuth)
	e.POST("/user/create", userController.Create)
	e.GET("/users/:id", userController.GetUser)
    e.GET("/users",userController.GetAllUsers)
	return e
}
