package repository

import "test-books-wishlist/internal/entity"

func (r *Repo) CreateUser(u *entity.User) error {
	err := r.SQLDB.Create(&u).Scan(&u).Error
	if err != nil {
		return err
	}
	return nil
}
