package service

import (
    "github.com/digitcodestudio/accounting-module/repository"
)

type ReportService struct {
    LedgerRepo repository.TransactionRepository
}

func NewReportService(repo repository.TransactionRepository) *ReportService {
    return &ReportService{LedgerRepo: repo}
}

// Neraca sederhana: return map[AccountID]saldo
func (s *ReportService) BalanceSheet() (map[string]float64, error) {
    txs, err := s.LedgerRepo.List()
    if err != nil {
        return nil, err
    }
    ledger := make(map[string]float64)
    for _, tx := range txs {
        ledger[tx.AccountID] += tx.Debit - tx.Credit
    }
    return ledger, nil
}

// Laba Rugi: hitung pendapatan dan beban
func (s *ReportService) IncomeStatement(accountTypes map[string]string) (income float64, expense float64, err error) {
    txs, err := s.LedgerRepo.List()
    if err != nil {
        return 0, 0, err
    }
    for _, tx := range txs {
        typ := accountTypes[tx.AccountID]
        switch typ {
		case "Pendapatan":
            income += tx.Credit - tx.Debit
        case "Beban":
            expense += tx.Debit - tx.Credit
        }
    }
    return income, expense, nil
}
