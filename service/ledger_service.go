package service

import (
    "github.com/digitcodestudio/accounting-module/repository"
)

type LedgerService struct {
    TransactionRepo repository.TransactionRepository
}

func NewLedgerService(txRepo repository.TransactionRepository) *LedgerService {
    return &LedgerService{TransactionRepo: txRepo}
}

// Simple Ledger: return map[AccountID]balance
func (s *LedgerService) GetLedger() (map[string]float64, error) {
    txs, err := s.TransactionRepo.List()
    if err != nil {
        return nil, err
    }

    ledger := make(map[string]float64)
    for _, tx := range txs {
        ledger[tx.AccountID] += tx.Debit - tx.Credit
    }
    return ledger, nil
}
