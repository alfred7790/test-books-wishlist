package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"test-books-wishlist/internal/entity"
	"testing"
)

func TestRepo_GetItem(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("failed to open sqlmock database: %v", err)
	}

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("failed to open gorm mock database: %v", err)
	}
	defer db.Close()

	expected := &entity.ItemWishList{
		WishListID: 1,
		BookID:     "XXXXXXXX",
		Book:       entity.Book{},
	}

	rows := sqlmock.NewRows([]string{"wish_list_id", "book_id"}).
		AddRow(expected.WishListID, expected.BookID)

	mock.ExpectQuery(`SELECT * FROM "item_wish_lists" WHERE (wish_list_id=$1 and book_id=$2) ORDER BY "item_wish_lists"."wish_list_id" ASC LIMIT 1`).
		WithArgs(expected.WishListID, expected.BookID).
		WillReturnRows(rows)

	sut := NewPostgresRepository()
	sut.SQLDB = gormDB
	sut.DBName = "test"
	actual, _, err := sut.GetItem(expected.WishListID, expected.BookID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, actual)
}
