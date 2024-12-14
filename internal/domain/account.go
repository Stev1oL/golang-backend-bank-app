package domain

import (
	"errors"
	"time"
)

type Account struct {
	ID            string    `json:"id"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Transaction struct {
	ID              string    `json:"id"`
	FromAccountID   string    `json:"from_account_id"`
	ToAccountID     string    `json:"to_account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"` // DEBIT or CREDIT
	QRCode          string    `json:"qr_code,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

type AccountRepository interface {
	Create(account *Account) error
	GetByID(id string) (*Account, error)
	GetByAccountNumber(accountNumber string) (*Account, error)
	Update(account *Account) error
	CreateTransaction(transaction *Transaction) error
	GetTransactions(accountID string) ([]Transaction, error)
}

type AccountService interface {
	CreateAccount(accountNumber string) (*Account, error)
	GetAccount(id string) (*Account, error)
	Transfer(fromAccountID, toAccountID string, amount float64) error
	GenerateQRCode(accountID string, amount float64) (string, error)
	ProcessQRPayment(qrCode string) error
	GetTransactionHistory(accountID string) ([]Transaction, error)
}

var (
	ErrInvalidAccountNumber = errors.New("invalid account number")
	ErrInsufficientBalance  = errors.New("insufficient balance")
	ErrAccountNotFound      = errors.New("account not found")
	ErrInvalidAmount        = errors.New("invalid amount")
)
