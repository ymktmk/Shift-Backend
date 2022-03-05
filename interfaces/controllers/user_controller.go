package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ymktmk/Shift-Backend/domain"
	"github.com/ymktmk/Shift-Backend/interfaces/database"
	"github.com/ymktmk/Shift-Backend/usecase"
	"gopkg.in/go-playground/validator.v9"
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

func (controller *UserController) Print(c echo.Context) (err error) {
	return c.JSON(http.StatusOK,c.Get("key"))
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// バリデーション
	if err = validator.New().Struct(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// 同じメールアドレス、uidでerr返ってくる → 同じものを挿入したときidは進む
	user, err := controller.Interactor.Add(*u)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := controller.Interactor.Update(*u)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}