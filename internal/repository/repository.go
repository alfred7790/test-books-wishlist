package repository

import "test-books-wishlist/internal/entity"

type Repo struct {
	Base
}

type Repository interface {
	Init(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error
	UsersRepo
	BooksRepo
}

type UsersRepo interface {
	CreateUser(user *entity.User) error
	UserAuth(username, password string) (*entity.User, error)
}

type BooksRepo interface {
	CreateWishList(wishlist *entity.WishList) (string, error)
	GetWishLists(userID uint) ([]*entity.WishList, string, error)
	AddItem(wishListID uint, book *entity.Book) (string, error)
}
