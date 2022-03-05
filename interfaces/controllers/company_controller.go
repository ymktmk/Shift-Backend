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

// 会社に所属している人を取得
func (controller *CompanyController) Show(c echo.Context) (err error) {
    // uidからその人の会社IDを取得 → company情報は取得しない
    user, err := controller.Interactor.UserByUid("xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
    // 会社IDからusers取得
    users, err := controller.Interactor.Users(user.CompanyID)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}