package usecase

import "github.com/ymktmk/Shift-Backend/domain"

type CompanyRepository interface {
	FindByUid(uid string) (user *domain.User, err error)
	FindUsersById(companyId int) (users *domain.Users, err error)
}

type CompanyInteractor struct {
    CompanyRepository CompanyRepository
}

func (interactor *CompanyInteractor) UserByUid(uid string) (user *domain.User, err error) {
    user, err = interactor.CompanyRepository.FindByUid(uid)
    return user, err
}

func (interactor *CompanyInteractor) Users(companyId int) (users *domain.Users, err error) {
    users, err = interactor.CompanyRepository.FindUsersById(companyId)
    return users, err
}