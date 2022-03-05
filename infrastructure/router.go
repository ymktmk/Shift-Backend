package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	companyController := controllers.NewCompanyController(NewSqlHandler())
	e := echo.New()
    g := e.Group("/admin")
	g.Use(verifyFirebaseToken)
	g.GET("/user",userController.Print)
	e.POST("/users/create", userController.Create)
	e.POST("/users/update", userController.Update)
	e.GET("/users/:id", userController.Show)
	e.GET("/company/users", companyController.Show)
	return e
}
