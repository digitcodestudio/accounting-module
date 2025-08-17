package models

import "time"

type Receivable struct {
    ID         string  `json:"id"`
    CustomerID string  `json:"customer_id"`
    Amount     float64 `json:"amount"`
    Received   float64 `json:"received"`
    DueDate    time.Time `json:"due_date"`
    CreatedAt  time.Time `json:"created_at"`
}
