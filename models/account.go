package models

type Account struct {
    ID       string `json:"id"`
    Code     string `json:"code"`
    Name     string `json:"name"`
    Type     string `json:"type"` // Harta, Hutang, Modal, Pendapatan, Beban
    ParentID string `json:"parent_id,omitempty"`
}
