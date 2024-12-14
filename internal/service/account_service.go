package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/steviol/golang-backend-bank-app/internal/domain"
)

type accountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) domain.AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount(accountNumber string) (*domain.Account, error) {
	if len(accountNumber) > 10 {
		return nil, domain.ErrInvalidAccountNumber
	}

	account := &domain.Account{
		ID:            uuid.New().String(),
		AccountNumber: accountNumber,
		Balance:       0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := s.repo.Create(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *accountService) GetAccount(id string) (*domain.Account, error) {
	return s.repo.GetByID(id)
}

func (s *accountService) AddBalance(accountNumber string, balance float64) error {
	if balance <= 0 {
		return domain.ErrInvalidAmount
	}

	account, err := s.repo.GetByAccountNumber(accountNumber)
	if err != nil {
		return err
	}

	account.Balance += balance
	account.UpdatedAt = time.Now()

	if err := s.repo.Update(account); err != nil {
		return err
	}

	return nil
}

func (s *accountService) Transfer(fromAccountID, toAccountID string, amount float64) error {
	if amount <= 0 {
		return domain.ErrInvalidAmount
	}

	fromAccount, err := s.repo.GetByID(fromAccountID)
	if err != nil {
		return err
	}

	toAccount, err := s.repo.GetByID(toAccountID)
	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return domain.ErrInsufficientBalance
	}

	// transfer
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	err = s.repo.Update(fromAccount)
	if err != nil {
		return nil
	}

	err = s.repo.Update(toAccount)
	if err != nil {
		return nil
	}

	// create record
	transaction := &domain.Transaction{
		ID:              uuid.New().String(),
		FromAccountID:   fromAccountID,
		ToAccountID:     toAccountID,
		Amount:          amount,
		TransactionType: "TRANSFER",
		CreatedAt:       time.Now(),
	}

	return s.repo.CreateTransaction(transaction)
}

func (s *accountService) GenerateQRCode(accountID string, amount float64) (string, error) {
	if amount < 0 {
		return "", domain.ErrInvalidAmount
	}

	qrData := struct {
		AccountID string  `json:"account_id"`
		Amount    float64 `json:"amount"`
		Timestamp int64   `json:"timestamp"`
	}{
		AccountID: accountID,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
	}

	jsonData, err := json.Marshal(qrData)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonData), nil
}

func (s *accountService) ProcessQRPayment(qrCode string) error {
	qrData, err := base64.StdEncoding.DecodeString(qrCode)
	if err != nil {
		return err
	}

	var data struct {
		AccountID string  `json:"account_id"`
		Amount    float64 `json:"amount"`
		Timestamp int64   `json:"timestamp"`
	}

	err = json.Unmarshal(qrData, &data)
	if err != nil {
		return err
	}

	if time.Now().Unix()-data.Timestamp > 300 {
		return fmt.Errorf("QR code expired")
	}

	return nil
}

func (s *accountService) GetTransactionHistory(accountID string) ([]domain.Transaction, error) {
	return s.repo.GetTransactions(accountID)
}
