// interfaces/databaseからのInputをusecase/user_repository.goで、
// interfaces/controllersへのGatewayをusecase/user_interactor.goで実現
// 継承してSQLを叩く外に依存していない
package usecase

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

type UserRepository interface {
	// SQLの実行メソッドをここに埋め込む
	FindByEmail(email string) (user domain.User, err error) 
	FindById(id int) (user domain.User, err error)
	FindAll() (users domain.Users, err error)
	Store(u domain.User) (user domain.User, err error)
	Update(u domain.User) (user domain.User, err error)
	DeleteById(user domain.User) (err error)
}