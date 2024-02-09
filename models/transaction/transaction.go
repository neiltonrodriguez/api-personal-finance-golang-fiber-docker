package transaction

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"personal-finance-api/database"
	"personal-finance-api/domain"
	AccountModel "personal-finance-api/models/account"
	TransactionTypeModel "personal-finance-api/models/transaction_type"
)

var Db *sql.DB

var (
	errTransactionNotFound = errors.New("transaction not found")
)

func Get(ctx context.Context, typeTransaction string) ([]domain.TransactionOutput, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	args := []interface{}{}
	conditions := ""
	if len(typeTransaction) > 0 {
		conditions = "WHERE type.title = ?"
		args = append(args, typeTransaction)
	}

	rows, err := Db.Query(`
	SELECT 
		t.id, 
		t.title, 
		t.value, 
		c.title, 
		type.title, 
		t.created_at, 
		t.updated_at 
	FROM golang.transaction t
	LEFT JOIN golang.category c ON t.category_id = c.id
	LEFT JOIN golang.transaction_type type ON t.type_transaction_id = type.id
	ORDER BY t.created_at DESC`+conditions, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.TransactionOutput
	for rows.Next() {
		var transaction domain.TransactionOutput
		err := rows.Scan(
			&transaction.Id,
			&transaction.Title,
			&transaction.Value,
			&transaction.Category,
			&transaction.TypeTransaction,
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func Create(ctx context.Context, t domain.TransactionInput) (domain.TransactionOutput, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.TransactionOutput{}, err
	}

	query := `INSERT INTO transaction (title, value, account_id, category_id, condition_id, type_transaction_id, type_payment_id, description, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	result, err := Db.ExecContext(ctx, query, t.Title, t.Value, t.AccountId, t.CategoryId, t.ConditionId, t.TypeTransactionId, t.TypePaymentId, t.Descritpion)
	if err != nil {
		return domain.TransactionOutput{}, err
	}
	defer Db.Close()

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	totalBalance, err := AccountModel.GetBalanceByAccount(ctx, t.AccountId)
	if err != nil {
		log.Fatal(err)
	}
	
	var newTotal float64
	typeTrans, err := TransactionTypeModel.GetById(ctx, t.TypeTransactionId)
	if typeTrans.Title == "gasto" {
		newTotal = totalBalance.Balance - t.Value
	} else {
		newTotal = totalBalance.Balance + t.Value
	}

	err = AccountModel.UpdateBalance(ctx, t.AccountId, newTotal)
	if err != nil {
		log.Fatal(err)
	}

	transaction, err := GetById(ctx, int(lastId))

	if err != nil {
		log.Fatal(err)
	}

	return transaction, nil
}

func Update(ctx context.Context, id int, u domain.User) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, u.Name, u.Email, u.Phone, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetById(ctx context.Context, id int) (domain.TransactionOutput, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.TransactionOutput{}, err
	}

	rows, err := Db.Query(`
	SELECT 
		t.id, 
		t.title, 
		t.value, 
		c.title, 
		co.title, 
		tt.title,
		p.title,
		COALESCE(t.description, ""),
		t.created_at, 
		t.updated_at 
	FROM golang.transaction t
	LEFT JOIN golang.category c ON t.category_id = c.id
	LEFT JOIN golang.transaction_type tt ON t.type_transaction_id = tt.id
	LEFT JOIN golang.payment_type p ON t.type_payment_id = p.id
	LEFT JOIN golang.condition co ON t.condition_id = co.id
	WHERE t.id = ? LIMIT 1`, id)
	if err != nil {
		return domain.TransactionOutput{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.TransactionOutput{}, errTransactionNotFound
	}
	var transaction domain.TransactionOutput
	err = rows.Scan(
		&transaction.Id,
		&transaction.Title,
		&transaction.Value,
		&transaction.Category,
		&transaction.Condition,
		&transaction.TypeTransaction,
		&transaction.TypePayment,
		&transaction.Descritpion,
		&transaction.CreatedAt,
		&transaction.UpdatedAt)
	if err != nil {
		return domain.TransactionOutput{}, err
	}

	return transaction, nil
}

func GetTransactionTotal(ctx context.Context) (domain.TransactionTotal, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.TransactionTotal{}, err
	}

	rows, err := Db.Query(`
	SELECT
	COALESCE(SUM(CASE WHEN t.type_transaction_id = 3 THEN t.value ELSE 0 END), 0) AS total_in,
	COALESCE(SUM(CASE WHEN t.type_transaction_id = 4 THEN t.value ELSE 0 END), 0) as total_out
	FROM golang.transaction t`)
	if err != nil {
		return domain.TransactionTotal{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.TransactionTotal{}, errTransactionNotFound
	}
	var transaction domain.TransactionTotal
	err = rows.Scan(
		&transaction.In,
		&transaction.Out,
	)
	if err != nil {
		return domain.TransactionTotal{}, err
	}

	return transaction, nil
}

func Delete(ctx context.Context, id int) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `DELETE FROM golang.users WHERE id = ?`

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
