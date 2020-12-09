package auth

import (
	"context"
	"errors"
	"movies/models/dto"
	"movies/models/entity/auth"
)

func AuthenticateUser(authToken string) (dto.User, error) {
	ctx := context.Background()
	var user auth.User
	err := user.Get(&ctx, authToken)
	if err != nil {
		return dto.User{}, errors.New("user not authorized")
	}
	return *dto.UserEntityToDto(user), nil
}
