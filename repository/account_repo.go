package repository

import "github.com/digitcodestudio/accounting-module/models"

type AccountRepository interface {
    Create(account *models.Account) error
    GetByID(id string) (*models.Account, error)
    List() ([]*models.Account, error)
}
