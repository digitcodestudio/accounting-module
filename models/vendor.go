package models

type Vendor struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Address string `json:"address"`
    Phone   string `json:"phone"`
    Email   string `json:"email"`
}
