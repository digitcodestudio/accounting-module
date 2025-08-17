package postgres

import (
    "database/sql"
    "github.com/digitcodestudio/accounting-module/models"
)

type PostgresAccountRepo struct {
    DB *sql.DB
}

func NewPostgresAccountRepo(db *sql.DB) *PostgresAccountRepo {
    return &PostgresAccountRepo{DB: db}
}

func (r *PostgresAccountRepo) Create(account *models.Account) error {
    _, err := r.DB.Exec(`INSERT INTO accounts (id, code, name, type, parent_id) VALUES ($1,$2,$3,$4,$5)`,
        account.ID, account.Code, account.Name, account.Type, account.ParentID)
    return err
}

func (r *PostgresAccountRepo) GetByID(id string) (*models.Account, error) {
    row := r.DB.QueryRow(`SELECT id, code, name, type, parent_id FROM accounts WHERE id=$1`, id)
    acc := &models.Account{}
    err := row.Scan(&acc.ID, &acc.Code, &acc.Name, &acc.Type, &acc.ParentID)
    if err != nil {
        return nil, err
    }
    return acc, nil
}

func (r *PostgresAccountRepo) List() ([]*models.Account, error) {
    rows, err := r.DB.Query(`SELECT id, code, name, type, parent_id FROM accounts`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var accounts []*models.Account
    for rows.Next() {
        acc := &models.Account{}
        rows.Scan(&acc.ID, &acc.Code, &acc.Name, &acc.Type, &acc.ParentID)
        accounts = append(accounts, acc)
    }
    return accounts, nil
}
