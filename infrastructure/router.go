package infrastructure

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/controllers"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

// c.Validateで使えるようになる
func (cv *CustomValidator) Validate(i interface{}) error {
	// cv.validator.RegisterValidation("email_check", Email)
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// func Email(fl validator.FieldLevel) bool {
// 	return fl.Field().String()
// }

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	companyController := controllers.NewCompanyController(NewSqlHandler())
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	// 認証なしapi
	e.POST("/api/user/create", userController.Create)
	// 認証つきapi
    g := e.Group("/api")
	g.Use(verifyFirebaseToken)
	g.GET("/user", userController.Show)
	g.PUT("/user/update", userController.Update)
	g.GET("/company/users", companyController.Show)
	return e
}
