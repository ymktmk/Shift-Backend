package usecase

import "github.com/ymktmk/Shift-Backend/domain"

type CompanyInteractor struct {
    CompanyRepository CompanyRepository
}

func (interactor *CompanyInteractor) UserByUid(uid string) (user domain.User, err error) {
    user, err = interactor.CompanyRepository.FindByUid(uid)
    return user, err
}

// 会社の人たち
func (interactor *CompanyInteractor) Users(companyId int) (users []domain.User, err error) {
    users, err = interactor.CompanyRepository.FindUsersById(companyId)
    return users, err
}