// ロジックをここに書く
package usecase

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u *domain.User) (user *domain.User, err error) {
    user, err = interactor.UserRepository.Store(u)
    if err != nil {
        return
    }
    return
}

func (interactor *UserInteractor) Update(uid string, u *domain.User) (user *domain.User, err error) {
    u, err = interactor.UserRepository.Update(uid, u)
    user, err = interactor.UserRepository.FindByUid(uid)
    return
}

func (interactor *UserInteractor) UserById(userId int) (user *domain.User, err error) {
    user, err = interactor.UserRepository.FindById(userId)
    return
}

func (interactor *UserInteractor) UserByUid(uid string) (user *domain.User, err error) {
    user, err = interactor.UserRepository.FindByUid(uid)
    return
}

func (interactor *UserInteractor) UserByEmail(email string) (user domain.User, err error) {
    user, err = interactor.UserRepository.FindByEmail(email)
    return
}