package usecase

import "github.com/ymktmk/Shift-Backend/domain"

type CompanyRepository interface {
	FindByUid(uid string) (user domain.User, err error)
	FindUsersById(companyId int) (users domain.Users, err error)
}