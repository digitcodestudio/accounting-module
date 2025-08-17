package postgres

import (
    "database/sql"
    "github.com/digitcodestudio/accounting-module/models"
)

type PostgresTransactionRepo struct {
    DB *sql.DB
}

func NewPostgresTransactionRepo(db *sql.DB) *PostgresTransactionRepo {
    return &PostgresTransactionRepo{DB: db}
}

func (r *PostgresTransactionRepo) Create(tx *models.Transaction) error {
    _, err := r.DB.Exec(
        `INSERT INTO transactions (id, account_id, debit, credit, description, created_at) 
         VALUES ($1,$2,$3,$4,$5,$6)`,
        tx.ID, tx.AccountID, tx.Debit, tx.Credit, tx.Desc, tx.CreatedAt,
    )
    return err
}

func (r *PostgresTransactionRepo) GetByID(id string) (*models.Transaction, error) {
    row := r.DB.QueryRow(
        `SELECT id, account_id, debit, credit, description, created_at 
         FROM transactions WHERE id=$1`, id,
    )
    tx := &models.Transaction{}
    err := row.Scan(&tx.ID, &tx.AccountID, &tx.Debit, &tx.Credit, &tx.Desc, &tx.CreatedAt)
    if err != nil {
        return nil, err
    }
    return tx, nil
}

func (r *PostgresTransactionRepo) List() ([]*models.Transaction, error) {
    rows, err := r.DB.Query(
        `SELECT id, account_id, debit, credit, description, created_at FROM transactions`,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var txs []*models.Transaction
    for rows.Next() {
        tx := &models.Transaction{}
        err := rows.Scan(&tx.ID, &tx.AccountID, &tx.Debit, &tx.Credit, &tx.Desc, &tx.CreatedAt)
        if err != nil {
            return nil, err
        }
        txs = append(txs, tx)
    }
    return txs, nil
}
