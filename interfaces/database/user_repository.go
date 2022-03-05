// User関連のSQL実行のコードを書く
package database

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

// ここで再定義して使う
type UserRepository struct {
    SqlHandler
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	// if err = repo.Find(&user, id).Error; err != nil {
	// 	return
	// }
	if err = repo.Joins("Company").Find(&user, id).Error; err != nil {
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

func (repo *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = repo.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) DeleteById(user domain.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}