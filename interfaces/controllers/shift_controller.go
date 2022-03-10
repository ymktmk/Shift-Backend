package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/interfaces/database"
	"github.com/ymktmk/Shift-Backend/usecase"
)

type ShiftController struct {
    Interactor usecase.ShiftInteractor
}

// 依存性を定義する
func NewShiftController(sqlHandler database.SqlHandler) *ShiftController {
    return &ShiftController{
        Interactor: usecase.ShiftInteractor{
            ShiftRepository: &database.ShiftRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *ShiftController) Print(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "Hello")
}