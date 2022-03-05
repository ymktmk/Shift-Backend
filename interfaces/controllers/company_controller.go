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

func (controller *CompanyController) Print(c echo.Context) (err error) {
	return c.JSON(http.StatusOK,"OK")
}