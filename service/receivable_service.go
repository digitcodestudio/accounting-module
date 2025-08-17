package service

import (
    "errors"
    "github.com/digitcodestudio/accounting-module/models"
    "github.com/digitcodestudio/accounting-module/repository"
    "time"
)

type ReceivableService struct {
    Repo repository.ReceivableRepository
}

func NewReceivableService(repo repository.ReceivableRepository) *ReceivableService {
    return &ReceivableService{Repo: repo}
}

func (s *ReceivableService) AddReceivable(r *models.Receivable) error {
    r.CreatedAt = time.Now()
    if r.Amount <= 0 {
        return errors.New("amount must be > 0")
    }
    return s.Repo.Create(r)
}

func (s *ReceivableService) Receive(id string, amount float64) error {
    r, err := s.Repo.GetByID(id)
    if err != nil {
        return err
    }
    r.Received += amount
    return s.Repo.Update(r)
}
