package entity

type Book struct {
	ID        string `json:"id" gorm:"primary_key" example:"KNYxEAAAQBAJ" descripcion:"GoogleBookID (ID) from google"`
	Authors   string `json:"authors" example:"Miranda De Moura" description:"Name of the author of the book"`
	Title     string `json:"title" example:"Viagens Na Madrugada" description:"Title of the book"`
	Publisher string `json:"publisher" example:"Clube de Autores" description:"Publisher of the book"`
}

type WishList struct {
	UserID uint   `json:"userId" gorm:"primary_key" example:"1" descripcion:"UserID"`
	BookID string `json:"bookId" gorm:"primary_key" example:"KNYxEAAAQBAJ" descripcion:"BookID"`
}
