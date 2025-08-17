package service

import (
    "errors"
    "github.com/digitcodestudio/accounting-module/models"
    "github.com/digitcodestudio/accounting-module/repository"
    "time"
)

type JournalService struct {
    AccountRepo     repository.AccountRepository
    TransactionRepo repository.TransactionRepository
}

func NewJournalService(accRepo repository.AccountRepository, txRepo repository.TransactionRepository) *JournalService {
    return &JournalService{
        AccountRepo:     accRepo,
        TransactionRepo: txRepo,
    }
}

func (s *JournalService) PostTransaction(tx *models.Transaction) error {
    account, err := s.AccountRepo.GetByID(tx.AccountID)
    if err != nil || account == nil {
        return errors.New("account not found")
    }

    tx.CreatedAt = time.Now()

    if err := s.TransactionRepo.Create(tx); err != nil {
        return err
    }

    return nil
}
