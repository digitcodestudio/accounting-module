package repository

import "github.com/digitcodestudio/accounting-module/models"

type PayableRepository interface {
    Create(payable *models.Payable) error
    GetByID(id string) (*models.Payable, error)
    List() ([]*models.Payable, error)
    Update(payable *models.Payable) error
}
