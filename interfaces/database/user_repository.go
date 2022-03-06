package database

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

type UserRepository struct {
    SqlHandler
}

func (repo *UserRepository) Store(u *domain.User) (user *domain.User, err error) {
	if err = repo.Create(u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u *domain.User) (user *domain.User, err error) {
	if err = repo.Save(u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) FindByUid(uid string) (user *domain.User, err error) {
	if err = repo.Joins("Company").Where("uid=?", uid).First(&user).Error; err != nil {
		return
	}
	return
}

// 使っていない

func (repo *UserRepository) FindById(userId int) (user *domain.User, err error) {
	if err = repo.Joins("Company").Find(&user, userId).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByEmail(email string) (user domain.User, err error) {
	if err = repo.Where("email=?", email).First(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) DeleteById(user domain.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}