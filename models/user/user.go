package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"personal-finance-api/database"
	"personal-finance-api/domain"
)

var Db *sql.DB

var (
	errUserNotFound = errors.New("user not found")
)

func Get(ctx context.Context) ([]domain.User, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		id, 
		name, 
		email, 
		phone, 
		COALESCE('address', ''), 
		COALESCE('zipcode', ''), 
		COALESCE('neighborhood', ''), 
		COALESCE('city', ''), 
		COALESCE('uf', ''), 
		COALESCE('number', ''), 
		created_at, 
		updated_at 
	FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Address,
			&user.Zipcode,
			&user.Neighborhood,
			&user.City,
			&user.UF,
			&user.Number,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func Create(ctx context.Context, u domain.User) (domain.User, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.User{}, err
	}

	password, err := domain.Hash(u.Password)
	if err != nil {
		return domain.User{}, err
	}

	query := `INSERT INTO users (name, email, password, phone, profile_id, created_at, updated_at) VALUES(?, ?, ?, ?, 2, NOW(), NOW())`

	result, err := Db.ExecContext(ctx, query, u.Name, u.Email, password, u.Phone)
	if err != nil {
		return domain.User{}, err
	}
	defer Db.Close()

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	user, err := GetById(ctx, int(lastId))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d", lastId)

	return user, nil
}

func Update(ctx context.Context, id int, u domain.User) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `
	UPDATE users 
	SET 
		name = ?, 
		email = ?, 
		phone = ?, 
		password = ?,
		address = ?,
		zipcode = ?,
		neighborhood = ?,
		city = ?,
		uf = ?,
		number = ?
	WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, u.Name, u.Email, u.Password, u.Phone, u.Address, u.Zipcode, u.Neighborhood, u.City, u.UF, u.Number, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetById(ctx context.Context, id int) (domain.User, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.User{}, err
	}

	rows, err := Db.Query(`
	SELECT
	    id, 
		name, 
		email, 
		phone, 
		COALESCE('address', ''), 
		COALESCE('zipcode', ''), 
		COALESCE('neighborhood', ''), 
		COALESCE('city', ''), 
		COALESCE('uf', ''), 
		COALESCE('number', ''), 
		created_at, 
		updated_at  
	FROM users WHERE id = ? limit 1`, id)
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
		&user.Phone,
		&user.Address,
		&user.Zipcode,
		&user.Neighborhood,
		&user.City,
		&user.UF,
		&user.Number,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func Delete(ctx context.Context, id int) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
