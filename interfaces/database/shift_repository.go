package database

import "github.com/ymktmk/Shift-Backend/domain"

type ShiftRepository struct {
	SqlHandler
}

func (repo *ShiftRepository) FindByUid(uid string) (user *domain.User, err error) {
	if err = repo.Where("uid=?", uid).First(&user).Error; err != nil {
		return
	}
	return
}