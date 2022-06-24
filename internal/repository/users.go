package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"test-books-wishlist/internal/entity"
)

func (r *Repo) CreateUser(u *entity.User) error {
	err := r.SQLDB.Create(&u).Scan(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UserAuth(username, password string) (*entity.User, error) {
	var user entity.User
	err := r.SQLDB.Where(&entity.User{UserName: username, Password: password}).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New(NotFound)
		}

		return nil, err
	}
	return &user, err
}
