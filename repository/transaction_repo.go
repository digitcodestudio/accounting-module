package repository

import "github.com/digitcodestudio/accounting-module/models"

type TransactionRepository interface {
    // Create simpan transaksi baru
    Create(tx *models.Transaction) error

    // GetByID ambil transaksi berdasarkan ID
    GetByID(id string) (*models.Transaction, error)

    // List ambil semua transaksi
    List() ([]*models.Transaction, error)
}
