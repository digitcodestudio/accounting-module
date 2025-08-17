package utils

import (
	"fmt"

	"github.com/digitcodestudio/accounting-module/models"
	"github.com/digitcodestudio/accounting-module/repository"
)

func AutoPost(tx *models.Transaction, txRepo repository.TransactionRepository) error {
    // Validasi sederhana
    if tx.Debit < 0 || tx.Credit < 0 {
        return ErrInvalidAmount
    }
    return txRepo.Create(tx)
}

var ErrInvalidAmount = fmt.Errorf("debit atau credit tidak boleh negatif")
