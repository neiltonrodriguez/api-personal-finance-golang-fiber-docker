package login

import (
	"context"
	"database/sql"
	"errors"
	"personal-finance-api/database"
	"personal-finance-api/domain"
)

var Db *sql.DB

var (
	errUserNotFound = errors.New("user not found")
)

func GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.User{}, err
	}

	rows, err := Db.Query(`SELECT id, name, email, password, phone, created_at, updated_at FROM golang.users WHERE email = ? limit 1`, email)
	if err != nil {
		return domain.User{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.User{}, errUserNotFound
	}
	var user domain.User
	err = rows.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
