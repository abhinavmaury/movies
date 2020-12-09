package auth

import (
	"context"
	"movies/base"
)

type User struct {
	UserId  int64
	Name    string
	Email   string
	IsAdmin int
}

func (user *User) Get(ctx *context.Context, authToken string) error {
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `SELECT u.user_id, u.username, u.email, u.is_admin
				FROM movies.users u
				INNER JOIN movies.user_auth a ON u.user_id = a.user_id
				WHERE a.auth_token = ?`
	row := conn.QueryRowContext(*ctx, query, authToken)
	if err = row.Scan(&user.UserId, &user.Name, &user.Email, &user.IsAdmin); err != nil {
		return err
	}
	return nil
}
