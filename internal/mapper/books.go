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
