package models

import "time"

type Payable struct {
    ID        string  `json:"id"`
    VendorID  string  `json:"vendor_id"`
    Amount    float64 `json:"amount"`
    Paid      float64 `json:"paid"`
    DueDate   time.Time `json:"due_date"`
    CreatedAt time.Time `json:"created_at"`
}
