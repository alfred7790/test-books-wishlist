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

/*func (r *Repo) CreateWishList(userID uint, books []*entity.Book) (string, error) {
	tx := r.SQLDB.Begin()

	defer tx.Rollback()

	for _, v := range books {
		err := tx.FirstOrCreate(v).Error
		if err != nil {
			return fmt.Sprintf("Could not save this book [%s:%s] at the moment", v.ID, v.Title), err
		}

		err = tx.FirstOrCreate(&entity.ItemWishList{
			UserID: userID,
			BookID: v.ID,
		}).Error
		if err != nil {
			return fmt.Sprintf("Could not save this book [%s:%s] at the moment as a wish list", v.ID, v.Title), err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return "Internal BD error, cannot save wishlist at the moment", err
	}

	return "", nil
}*/

func (r *Repo) GetWishLists(userID uint) ([]*entity.WishList, string, error) {
	var lists []*entity.WishList
	err := r.SQLDB.Model(&entity.WishList{ID: userID}).Preload("Items.Book").Find(&lists).Error
	if err != nil {
		return lists, "Cannot get wishlist at the moment", err
	}

	return lists, "", nil
}
