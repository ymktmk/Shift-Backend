package controllers

import (
	"net/http"

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

func (controller *UserController) Create(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = validator.New().Struct(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// 同じメールアドレス、uidでerr返ってくる → 同じものを挿入したときidは進む
	user, err := controller.Interactor.Add(u)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c echo.Context) (err error) {
	uid := c.Get("uid").(string)
	// Jsonから構造体に変換
	uur := new(domain.UserUpdateRequest)
	if err = c.Bind(uur); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// バリデーション
	if err = validator.New().Struct(uur); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// DTOをUserのEntityに変換
	u := domain.User{Name: uur.Name}
	user, err := controller.Interactor.Update(uid, &u)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// レスポンス
	// resopnse := domain.UserUpdateResponse{UID: user.UID, Name: user.Name, Email: user.Email}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Show(c echo.Context) (err error) {
	uid := c.Get("uid").(string)
	user, err := controller.Interactor.UserByUid(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}