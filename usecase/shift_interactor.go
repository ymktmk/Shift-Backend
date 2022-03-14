package usecase

import "github.com/ymktmk/Shift-Backend/domain"

type ShiftRepository interface {
	FindByUid(uid string) (user *domain.User, err error)
}

type ShiftInteractor struct {
    ShiftRepository ShiftRepository
}

func (interactor *ShiftInteractor) UserByUid(uid string) (user *domain.User, err error) {
    user, err = interactor.ShiftRepository.FindByUid(uid)
    return
}