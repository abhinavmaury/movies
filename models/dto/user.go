package dto

import "movies/models/entity/auth"

type User struct {
	UserId  int64  `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin int    `json:"is_admin"`
	UserAuth
}

type UserAuth struct {
	AuthToken string `json:"auth_token"`
}

func UserEntityToDto(user auth.User) *User {
	return &User{
		UserId:  user.UserId,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}