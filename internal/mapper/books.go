package mapper

import (
	"strings"
	"test-books-wishlist/internal/entity"
)

func GoogleBookToBooksDTO(b *entity.GoogleBooksResponse) *entity.BooksDTO {
	respose := entity.BooksDTO{
		Results: b.TotalItems,
		Items:   len(b.Items),
		Books:   make([]*entity.Book, 0),
	}

	for _, v := range b.Items {
		respose.Books = append(respose.Books, &entity.Book{
			ID:        v.ID,
			Authors:   strings.Join(v.VolumeInfo.Authors, "|"),
			Title:     v.VolumeInfo.Title,
			Publisher: v.VolumeInfo.Publisher,
		})
	}

	return &respose
}

func WishListMapper(lists []*entity.WishList) []*entity.WishListDTO {
	response := make([]*entity.WishListDTO, 0)
	for _, v := range lists {
		response = append(response, &entity.WishListDTO{
			ID:        v.ID,
			UserID:    v.UserID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			Items:     BooksMapper(v.Items),
		})
	}

	return response
}

func BooksMapper(books []*entity.ItemWishList) []*entity.Book {
	response := make([]*entity.Book, 0)
	for _, v := range books {
		response = append(response, &entity.Book{
			ID:        v.BookID,
			Authors:   v.Book.Authors,
			Title:     v.Book.Title,
			Publisher: v.Book.Publisher,
		})
	}

	return response
}
