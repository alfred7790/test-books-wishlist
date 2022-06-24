package repository

import (
	"fmt"
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
