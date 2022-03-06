package database

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

type CompanyRepository struct {
    SqlHandler
}

func (repo *CompanyRepository) FindByUid(uid string) (user *domain.User, err error) {
	if err = repo.Where("uid=?", uid).First(&user).Error; err != nil {
		return
	}
	return
}

func (repo *CompanyRepository) FindUsersById(companyId int) (users *domain.Users, err error) {
	if err = repo.Joins("Company").Where("company_id=?", companyId).Find(&users).Error; err != nil {
		return
	}
	return
}