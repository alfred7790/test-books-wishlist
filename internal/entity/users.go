package entity

type User struct {
	Id       uint   `json:"id" gorm:"primary_key" example:"1" descripcion:"Primary key"`
	UserName string `json:"username" gorm:"not null;unique" example:"Pantufla89" description:"Username"`
	Password string `json:"password" gorm:"not null" example:"Pantufla1234" description:"Password"`
}
