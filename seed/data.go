package seed

import (
	"github.com/spootrick/survi/api/model"
)

var users = []model.User{
	{
		FirstName: "Furkan",
		LastName:  "Karakoyunlu",
		Email:     "furkan@karakoyunlu.com",
		Password:  "123456",
		Roles:     "ROLE_ADMIN,ROLE_USER",
	},
}
