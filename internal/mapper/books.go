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

func WishListsDTOMapper(lists []*entity.WishList) []*entity.WishListDTO {
	response := make([]*entity.WishListDTO, 0)
	for _, v := range lists {
		response = append(response, WishListDTOMapper(v))
	}

	return response
}

func WishListDTOMapper(list *entity.WishList) *entity.WishListDTO {
	return &entity.WishListDTO{
		ID:        list.ID,
		UserID:    list.UserID,
		Name:      list.Name,
		CreatedAt: list.CreatedAt,
		Items:     BooksMapper(list.Items),
	}
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
