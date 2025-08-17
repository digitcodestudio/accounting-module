package service

import (
    "errors"
    "github.com/digitcodestudio/accounting-module/models"
    "github.com/digitcodestudio/accounting-module/repository"
    "time"
)

type PayableService struct {
    Repo repository.PayableRepository
}

func NewPayableService(repo repository.PayableRepository) *PayableService {
    return &PayableService{Repo: repo}
}

func (s *PayableService) AddPayable(p *models.Payable) error {
    p.CreatedAt = time.Now()
    if p.Amount <= 0 {
        return errors.New("amount must be > 0")
    }
    return s.Repo.Create(p)
}

func (s *PayableService) Pay(id string, amount float64) error {
    p, err := s.Repo.GetByID(id)
    if err != nil {
        return err
    }
    p.Paid += amount
    return s.Repo.Update(p)
}
