package transaction

import (
	"context"
	"database/sql"
	"personal-finance-api/database"
	"personal-finance-api/domain"
)

var Db *sql.DB

func Get(ctx context.Context) (domain.TransactionTypes, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		t.id, 
		t.title, 
		COALESCE(t.description, '')  
	FROM transaction_type t`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactionsTypes domain.TransactionTypes
	for rows.Next() {
		var transactionType domain.TransactionType
		err := rows.Scan(
			&transactionType.Id,
			&transactionType.Title,
			&transactionType.Descritpion)
		if err != nil {
			return nil, err
		}
		transactionsTypes = append(transactionsTypes, transactionType)
	}

	return transactionsTypes, nil
}

// func Create(ctx context.Context, t domain.TransactionInput) (domain.TransactionOutput, error) {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return domain.TransactionOutput{}, err
// 	}

// 	query := `INSERT INTO transaction (title, value, account_id, category_id, condition_id, type_transaction_id, type_payment_id, description, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

// 	result, err := Db.ExecContext(ctx, query, t.Title, t.Value, t.AccountId, t.CategoryId, t.ConditionId, t.TypeTransactionId, t.TypePaymentId, t.Descritpion)
// 	if err != nil {
// 		return domain.TransactionOutput{}, err
// 	}
// 	defer Db.Close()

// 	lastId, err := result.LastInsertId()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	totalBalance, err := AccountModel.GetBalanceByAccount(ctx, t.AccountId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var newTotal float64
// 	if t.TypeTransactionId == 1 {
// 		newTotal = totalBalance.Balance - t.Value
// 	} else {
// 		newTotal = totalBalance.Balance + t.Value
// 	}

// 	err = AccountModel.UpdateBalance(ctx, t.AccountId, newTotal)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	transaction, err := GetById(ctx, int(lastId))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return transaction, nil
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

// func GetById(ctx context.Context, id int) (domain.TransactionOutput, error) {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return domain.TransactionOutput{}, err
// 	}

// 	rows, err := Db.Query(`
// 	SELECT
// 		t.id,
// 		t.title,
// 		t.value,
// 		c.title,
// 		co.title,
// 		tt.title,
// 		p.title,
// 		COALESCE(t.description, ""),
// 		t.created_at,
// 		t.updated_at
// 	FROM finance.transaction t
// 	LEFT JOIN finance.category c ON t.category_id = c.id
// 	LEFT JOIN finance.transaction_type tt ON t.type_transaction_id = tt.id
// 	LEFT JOIN finance.payment_type p ON t.type_payment_id = p.id
// 	LEFT JOIN finance.condition co ON t.condition_id = co.id
// 	WHERE t.id = ? LIMIT 1`, id)
// 	if err != nil {
// 		return domain.TransactionOutput{}, err
// 	}

// 	defer rows.Close()

// 	rowExist := rows.Next()
// 	if !rowExist {
// 		return domain.TransactionOutput{}, errTransactionNotFound
// 	}
// 	var transaction domain.TransactionOutput
// 	err = rows.Scan(
// 		&transaction.Id,
// 		&transaction.Title,
// 		&transaction.Value,
// 		&transaction.Category,
// 		&transaction.Condition,
// 		&transaction.TypeTransaction,
// 		&transaction.TypePayment,
// 		&transaction.Descritpion,
// 		&transaction.CreatedAt,
// 		&transaction.UpdatedAt)
// 	if err != nil {
// 		return domain.TransactionOutput{}, err
// 	}

// 	return transaction, nil
// }

// func GetTransactionTotal(ctx context.Context) (domain.TransactionTotal, error) {
// 	var err error
// 	Db, err = database.ConnectToDB()
// 	if err != nil {
// 		return domain.TransactionTotal{}, err
// 	}

// 	rows, err := Db.Query(`
// 	SELECT
// 	SUM(CASE WHEN t.type_transaction_id = 2 THEN t.value ELSE 0 END) AS total_in,
// 	SUM(CASE WHEN t.type_transaction_id = 1 THEN t.value ELSE 0 END) AS total_out
// 	FROM transaction t`)
// 	if err != nil {
// 		return domain.TransactionTotal{}, err
// 	}

// 	defer rows.Close()

// 	rowExist := rows.Next()
// 	if !rowExist {
// 		return domain.TransactionTotal{}, errTransactionNotFound
// 	}
// 	var transaction domain.TransactionTotal
// 	err = rows.Scan(
// 		&transaction.In,
// 		&transaction.Out,
// 	)
// 	if err != nil {
// 		return domain.TransactionTotal{}, err
// 	}

// 	return transaction, nil
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
