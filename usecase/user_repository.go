// interfaces/databaseからのInputをusecase/user_repository.goで、
// interfaces/controllersへのGatewayをusecase/user_interactor.goで実現
// 継承してSQLを叩く外に依存していない
package usecase

import (
	"github.com/ymktmk/Shift-Backend/domain"
)

type UserRepository interface {
	// SQLの実行メソッドをここに埋め込む
	Store(user domain.User) (int, error)
	FindById(identifier int) (domain.User, error)
	FindAll() (domain.Users, error)
}