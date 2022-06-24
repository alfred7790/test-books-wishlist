package entity

import "time"

type Book struct {
	ID        string `json:"id" gorm:"primary_key" example:"KNYxEAAAQBAJ" descripcion:"GoogleBookID (ID) from google"`
	Authors   string `json:"authors" example:"Miranda De Moura" description:"Name of the author of the book"`
	Title     string `json:"title" example:"Viagens Na Madrugada" description:"Title of the book"`
	Publisher string `json:"publisher" example:"Clube de Autores" description:"Publisher of the book"`
}

type WishList struct {
	ID        uint            `json:"-" gorm:"primary_key"`
	UserID    uint            `json:"-" gorm:"not null"`
	Name      string          `json:"name" gorm:"not null" example:"MyFirstWishList" descripcion:"wishlist's name'"`
	CreatedAt time.Time       `json:"-" gorm:"default:current_timestamp"`
	Items     []*ItemWishList `json:"-" gorm:"foreigner:WishListID"`
}

type ItemWishList struct {
	WishListID uint   `json:"wishListId" gorm:"primary_key" example:"1" descripcion:"WishListID"`
	BookID     string `json:"bookId" gorm:"primary_key" example:"KNYxEAAAQBAJ" descripcion:"BookID"`
	Book       Book   `json:"book" gorm:"foreigner:BookID"`
}

type BooksDTO struct {
	Results int     `json:"results" example:"999" descripcion:"Total of results"`
	Items   int     `json:"items" example:"20" descripcion:"items in this page"`
	Books   []*Book `json:"books"`
}

type WishListDTO struct {
	ID        uint      `json:"id" example:"1" descripcion:"wishlist identifier"`
	UserID    uint      `json:"userId" example:"1" descripcion:"user identifier"`
	Name      string    `json:"name" example:"MyFirstWishList" descripcion:"wishlist's name'"`
	CreatedAt time.Time `json:"createdAt" example:"2022-06-24T19:38:46.814728Z" descripcion:"date of creation"`
	Items     []*Book   `json:"books"`
}
