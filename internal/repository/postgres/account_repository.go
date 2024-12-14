package postgres

import (
	"database/sql"
	"time"

	"github.com/steviol/golang-backend-bank-app/internal/domain"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *domain.Account) error {
	query := `
		INSERT INTO accounts (id, account_number, balance, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query,
		account.ID,
		account.AccountNumber,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	return err
}

func (r *accountRepository) GetByID(id string) (*domain.Account, error) {
	account := &domain.Account{}
	query := `
		SELECT id, account_number, balance, created_at, updated_at
		FROM accounts WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	return account, err
}

func (r *accountRepository) GetByAccountNumber(accountNumber string) (*domain.Account, error) {
	account := &domain.Account{}
	query := `
		SELECT id, account_number, balance, created_at, updated_at
		FROM accounts WHERE account_number = $1
	`
	err := r.db.QueryRow(query, accountNumber).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	return account, err
}

func (r *accountRepository) Update(account *domain.Account) error {
	query := `
		UPDATE accounts
		SET balance = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(query,
		account.Balance,
		time.Now(),
		account.ID,
	)
	return err
}

func (r *accountRepository) CreateTransaction(transaction *domain.Transaction) error {
	query := `
		INSERT INTO transactions (id, from_account_id, to_account_id, amount, transaction_type, qr_code, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query,
		transaction.ID,
		transaction.FromAccountID,
		transaction.ToAccountID,
		transaction.Amount,
		transaction.TransactionType,
		transaction.QRCode,
		transaction.CreatedAt,
	)
	return err
}

func (r *accountRepository) GetTransactions(accountID string) ([]domain.Transaction, error) {
	query := `
		SELECT id, from_account_id, to_account_id, amount, transaction_type, qr_code, created_at
		FROM transactions
		WHERE from_account_id = $1 OR to_account_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		err := rows.Scan(
			&t.ID,
			&t.FromAccountID,
			&t.ToAccountID,
			&t.Amount,
			&t.TransactionType,
			&t.QRCode,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
