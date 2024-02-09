package category

import (
	"context"
	"database/sql"
	"personal-finance-api/database"
	"personal-finance-api/domain"
)

var Db *sql.DB

func Get(ctx context.Context) (domain.Categorys, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		c.id, 
		c.title, 
		COALESCE(c.description, '')
	FROM golang.category c`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorys domain.Categorys
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(
			&category.Id,
			&category.Title,
			&category.Description)
		if err != nil {
			return nil, err
		}
		categorys = append(categorys, category)
	}

	return categorys, nil
}

// func Create(ctx context.Context, a domain.AccountInput) (domain.AccountOutput, error) {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return domain.AccountOutput{}, err
// 	}

// 	query := `INSERT INTO account (title, balance, type_account, bank_id, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`

// 	result, err := Db.ExecContext(ctx, query, a.Title, a.Balance, a.TypeAccount, a.BankId)
// 	if err != nil {
// 		return domain.AccountOutput{}, err
// 	}
// 	defer Db.Close()

// 	lastId, err := result.LastInsertId()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	account, err := GetById(ctx, int(lastId))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return account, nil
// }

// func UpdateBalance(ctx context.Context, accountId int, value float64) error {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return err
// 	}

// 	query := `UPDATE account SET balance = ? WHERE id = ?`

// 	_, err = Db.ExecContext(ctx, query, value, accountId)
// 	if err != nil {
// 		return err
// 	}
// 	defer Db.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return nil
// }

// func GetById(ctx context.Context, id int) (domain.AccountOutput, error) {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return domain.AccountOutput{}, err
// 	}

// 	rows, err := Db.Query(`
// 	SELECT
// 		a.id,
// 		a.title,
// 		a.balance,
// 		t.title,
// 		b.title,
// 		created_at,
// 		updated_at
// 	FROM account a
// 	LEFT JOIN type_account t ON a.type_account = t.id
// 	LEFT JOIN bank b ON a.bank_id = b.id
// 	WHERE a.id = ?`, id)
// 	if err != nil {
// 		return domain.AccountOutput{}, err
// 	}

// 	defer rows.Close()

// 	rowExist := rows.Next()
// 	if !rowExist {
// 		return domain.AccountOutput{}, errAccountNotFound
// 	}
// 	var account domain.AccountOutput
// 	err = rows.Scan(
// 		&account.Id,
// 		&account.Title,
// 		&account.Balance,
// 		&account.TypeAccount,
// 		&account.Bank,
// 		&account.CreatedAt,
// 		&account.UpdatedAt)
// 	if err != nil {
// 		return domain.AccountOutput{}, err
// 	}

// 	return account, nil
// }

// func Update(ctx context.Context, id int, u domain.User) error {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return err
// 	}

// 	query := `UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?`

// 	_, err = Db.ExecContext(ctx, query, u.Name, u.Email, u.Phone, id)
// 	if err != nil {
// 		return err
// 	}
// 	defer Db.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return nil
// }

// func Delete(ctx context.Context, id int) error {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return err
// 	}

// 	query := `DELETE FROM users WHERE id = ?`

// 	_, err = Db.ExecContext(ctx, query, id)
// 	if err != nil {
// 		return err
// 	}
// 	defer Db.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return nil
// }
