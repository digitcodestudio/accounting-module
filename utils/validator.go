package utils

import "github.com/digitcodestudio/accounting-module/models"

func ValidateTransaction(tx *models.Transaction) bool {
	return tx.Debit >= 0 && tx.Credit >= 0
}