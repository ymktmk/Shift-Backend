package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
    e := echo.New()
    // g := e.Group("/admin")
	// g.Use(FirebaseAuth)
	e.POST("/users/create", userController.Create)
	// e.POST("/users/update", userController.Update)
	e.GET("/users/:id", userController.GetUser)
	return e
}
