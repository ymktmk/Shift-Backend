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

// シフトを作成する
func (controller *ShiftController) Create(c echo.Context) (err error) {
	// uidを取得
	uid := c.Get("uid").(string)
	// userを取得する
	user, err := controller.Interactor.UserByUid(uid)
	

	return c.JSON(http.StatusOK, user)
}