package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"test-books-wishlist/internal/entity"
)

func (r *Repo) CreateWishList(wishlist *entity.WishList) (string, error) {
	err := r.SQLDB.Create(wishlist).Scan(wishlist).Error
	if err != nil {
		return fmt.Sprintf("Could not save this wishlist [%s] at the moment", wishlist.Name), err
	}

	return "", nil
}

func (r *Repo) GetWishLists(userID uint) ([]*entity.WishList, string, error) {
	var lists []*entity.WishList
	err := r.SQLDB.Model(&entity.WishList{ID: userID}).Preload("Items.Book").Find(&lists).Error
	if err != nil {
		return lists, "Cannot get wishlist at the moment", err
	}

	return lists, "", nil
}

func (r *Repo) AddItem(wishListID uint, book *entity.Book) (string, error) {
	tx := r.SQLDB.Begin()

	defer tx.Rollback()

	err := tx.FirstOrCreate(book).Error
	if err != nil {
		return fmt.Sprintf("Could not save this book [%s:%s] at the moment", book.ID, book.Title), err
	}

	err = tx.FirstOrCreate(&entity.ItemWishList{
		WishListID: wishListID,
		BookID:     book.ID,
	}).Error
	if err != nil {
		return fmt.Sprintf("Could not save this book [%s:%s] at the moment as a wish list", book.ID, book.Title), err
	}

	if err := tx.Commit().Error; err != nil {
		return "Internal BD error, cannot save wishlist at the moment", err
	}

	return "", nil
}

func (r *Repo) RemoveItem(wishListID uint, bookId string) (string, error) {
	err := r.SQLDB.Where("wish_list_id=? and book_id=?", wishListID, bookId).Delete(&entity.ItemWishList{}).Error
	if err != nil {
		return "Cannot delete this book at the moment", err
	}

	return "", nil
}

func (r *Repo) GetItem(wishListID uint, bookId string) (*entity.ItemWishList, string, error) {
	var found entity.ItemWishList
	err := r.SQLDB.Model(entity.ItemWishList{}).
		Where("wish_list_id=? and book_id=?", wishListID, bookId).First(&found).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, "This book doesn't exist", err
		}
		return nil, "Cannot get this book at the moment", err
	}

	return &found, "", nil
}

func (r *Repo) GetBooksByWishList(wishlistID uint) (*entity.WishList, string, error) {
	var list entity.WishList
	err := r.SQLDB.Model(&entity.WishList{ID: wishlistID}).Preload("Items.Book").First(&list).Error
	if err != nil {
		return &list, "Cannot get wishlist at the moment", err
	}

	return &list, "", nil
}
