package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
	"gopkg.in/go-playground/validator.v9"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	companyController := controllers.NewCompanyController(NewSqlHandler())
	shiftController := controllers.NewShiftController(NewSqlHandler())
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	// 認証なしapi
	e.POST("/api/user/create", userController.Create)
	// 認証つきapi
	g := e.Group("/api")
	g.Use(verifyFirebaseToken)
	g.GET("/user", userController.Show)
	g.PUT("/user/update", userController.Update)
	g.GET("/company/users", companyController.Show)
	// shift
	g.POST("/shift/create", shiftController.Create)
	return e
}
