package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/domain"
	"github.com/ymktmk/Shift-Backend/interfaces/database"
	"github.com/ymktmk/Shift-Backend/usecase"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// バリデーション
	// ユーザー登録処理 → 同じメールアドレスでerr返ってくる
	user, err := controller.Interactor.Add(*u)
    if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}