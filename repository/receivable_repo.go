package repository

import "github.com/digitcodestudio/accounting-module/models"

type ReceivableRepository interface {
    Create(receivable *models.Receivable) error
    GetByID(id string) (*models.Receivable, error)
    List() ([]*models.Receivable, error)
    Update(receivable *models.Receivable) error
}
