package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	companyController := controllers.NewCompanyController(NewSqlHandler())
	e := echo.New()
	e.POST("/user/create", userController.Create)
    g := e.Group("/api")
	g.Use(verifyFirebaseToken)
	
	e.PUT("/user/update", userController.Update)
	// 認証つきapi
	g.GET("/user", userController.Show)
	g.GET("/company/users", companyController.Show)
	return e
}
