package mocks

import (
	"test-books-wishlist/internal/entity"
	"test-books-wishlist/internal/repository"
)

type RepositoryMock struct {
	repository.BooksRepo
	repository.UsersRepo

	InitMockFn func(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error

	CreateWishListMockFn     func(wishlist *entity.WishList) (string, error)
	GetWishListsMockFn       func(userID uint) ([]*entity.WishList, string, error)
	AddItemMockFn            func(wishListID uint, book *entity.Book) (string, error)
	RemoveItemMockFn         func(wishListID uint, bookId string) (string, error)
	GetItemMockFn            func(wishListID uint, bookId string) (*entity.ItemWishList, string, error)
	GetBooksByWishListMockFn func(wishlistID uint) (*entity.WishList, string, error)
	RemoveWishListMockFn     func(wishListID uint) (string, error)
	WishListExistsMockFn     func(wishListID uint) (bool, string, error)

	CreateUserMockFn func(user *entity.User) error
	UserAuthMockFn   func(username, password string) (*entity.User, error)
}

func (m *RepositoryMock) Init(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error {
	return m.InitMockFn(dbIP, dbPort, dbUser, dbPass, dbName, retryCount)
}

func (m *RepositoryMock) CreateWishList(wishlist *entity.WishList) (string, error) {
	return m.CreateWishListMockFn(wishlist)
}
func (m *RepositoryMock) GetWishLists(userID uint) ([]*entity.WishList, string, error) {
	return m.GetWishListsMockFn(userID)
}
func (m *RepositoryMock) AddItem(wishListID uint, book *entity.Book) (string, error) {
	return m.AddItemMockFn(wishListID, book)
}
func (m *RepositoryMock) RemoveItem(wishListID uint, bookId string) (string, error) {
	return m.RemoveItemMockFn(wishListID, bookId)
}
func (m *RepositoryMock) GetItem(wishListID uint, bookId string) (*entity.ItemWishList, string, error) {
	return m.GetItemMockFn(wishListID, bookId)
}
func (m *RepositoryMock) GetBooksByWishList(wishlistID uint) (*entity.WishList, string, error) {
	return m.GetBooksByWishListMockFn(wishlistID)
}
func (m *RepositoryMock) RemoveWishList(wishListID uint) (string, error) {
	return m.RemoveWishListMockFn(wishListID)
}
func (m *RepositoryMock) WishListExists(wishListID uint) (bool, string, error) {
	return m.WishListExistsMockFn(wishListID)
}

func (m *RepositoryMock) CreateUser(user *entity.User) error {
	return m.CreateUserMockFn(user)
}
func (m *RepositoryMock) UserAuth(username, password string) (*entity.User, error) {
	return m.UserAuthMockFn(username, password)
}
