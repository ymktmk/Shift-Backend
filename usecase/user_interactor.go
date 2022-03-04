// ロジックをここに書く
package usecase

import (
	"github.com/ymktmk/Shift-Backend/domain"
	// "golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
    u, err := interactor.UserRepository.Store(u)
    user, err := interactor.UserRepository.FindById(int(u.Model.ID))
    return user, err
}

func (interactor *UserInteractor) Update(u domain.User) (domain.User, error) {
    u, err := interactor.UserRepository.Update(u)
    user, err := interactor.UserRepository.FindById(int(u.Model.ID))
    return user, err
}

func (interactor *UserInteractor) UserById(id int) (domain.User, error) {
    user, err := interactor.UserRepository.FindById(id)
    return user, err
}

func (interactor *UserInteractor) UserByUid(uid string) (domain.User, error) {
    user, err := interactor.UserRepository.FindByUid(uid)
    return user, err
}

func (interactor *UserInteractor) UserByEmail(email string) (domain.User, error) {
    user, err := interactor.UserRepository.FindByEmail(email)
    return user, err
}