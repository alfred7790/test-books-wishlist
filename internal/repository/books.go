package repository

import (
	"fmt"
	"test-books-wishlist/internal/entity"
)

func (r *Repo) CreateWishList(userID uint, books []*entity.Book) (string, error) {
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
}
