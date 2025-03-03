package models

type Product struct {
	ID          int     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}
