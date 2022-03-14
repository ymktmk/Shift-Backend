package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/database"
	"github.com/ymktmk/Shift-Backend/usecase"
)

type CompanyController struct {
    Interactor usecase.CompanyInteractor
}

func NewCompanyController(sqlHandler database.SqlHandler) *CompanyController {
    return &CompanyController{
        Interactor: usecase.CompanyInteractor{
            CompanyRepository: &database.CompanyRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

// 会社の人を取得する
func (controller *CompanyController) Show(c echo.Context) (err error) {
    uid := c.Get("uid").(string)
    user, err := controller.Interactor.UserByUid(uid)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
    users, err := controller.Interactor.Users(user.CompanyID)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}