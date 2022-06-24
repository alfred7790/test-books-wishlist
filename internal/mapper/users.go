package mapper

import "test-books-wishlist/internal/entity"

func UserDTOMapper(u *entity.User) entity.UserDTO {
	return entity.UserDTO{
		Id:       u.Id,
		UserName: u.UserName,
	}
}
