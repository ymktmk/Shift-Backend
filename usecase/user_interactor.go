// ロジックをここに書く
package usecase

import (
	"github.com/ymktmk/Shift-Backend/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
    UserRepository UserRepository
}

// ユーザー作成
func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
    // 同じメールアドレスがいないかチェック
    // パスワードハッシュ化
    hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    u.Password = string(hashPassword)
    // メールアドレス同じ人がいたらエラー
    u, err := interactor.UserRepository.Store(u)
    user, err := interactor.UserRepository.FindById(u.ID)
    return user, err
}

func (interactor *UserInteractor) Users() (domain.Users, error) {
    users, err := interactor.UserRepository.FindAll()
    return users, err
}

func (interactor *UserInteractor) UserById(id int) (domain.User, error) {
    user, err := interactor.UserRepository.FindById(id)
    return user, err
}

func (interactor *UserInteractor) UserByEmail(email string) (domain.User, error) {
    user, err := interactor.UserRepository.FindByEmail(email)
    return user, err
}