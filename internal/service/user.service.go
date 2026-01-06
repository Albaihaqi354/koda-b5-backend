package service

import "github.com/Albaihaqi354/koda-b5-backend/internal/dto"

var Users []dto.User

func SaveUser(user dto.User) {
	Users = append(Users, user)
}

func FindUser(email, password string) *dto.User {
	for _, user := range Users {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}
