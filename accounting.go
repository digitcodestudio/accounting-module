package accounting

import (
    "github.com/digitcodestudio/accounting-module/repository"
    "github.com/digitcodestudio/accounting-module/service"
)

type Accounting struct {
    Journal *service.JournalService
}

func NewAccounting(accRepo repository.AccountRepository, txRepo repository.TransactionRepository) *Accounting {
    journal := service.NewJournalService(accRepo, txRepo)
    return &Accounting{
        Journal: journal,
    }
}
