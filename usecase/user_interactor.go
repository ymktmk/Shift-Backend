// ロジックをここに書く
package usecase

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
    u, err := interactor.UserRepository.Store(u)
    user, err := interactor.UserRepository.FindById(u.ID)
    return user, err
}

func (interactor *UserInteractor) Users() (domain.Users, error) {
    users, err := interactor.UserRepository.FindAll()
    return users, err
}

func (interactor *UserInteractor) UserById(id int) (domain.User, error) {
    user, err := interactor.UserRepository.FindById(id)
    return user, err
}