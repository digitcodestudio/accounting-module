package models

import "time"

type Transaction struct {
    ID        string  `json:"id"`
    AccountID string  `json:"account_id"`
    Debit     float64 `json:"debit"`
    Credit    float64 `json:"credit"`
    Desc      string  `json:"description"`
    CreatedAt time.Time
}
