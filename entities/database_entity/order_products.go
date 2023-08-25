package database_entity

import "time"

type OrderProducts struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	OrderId   int        `json:"order_id"`
	ProductId int        `json:"product_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
	DeletedBy string     `json:"deleted_by"`
}

func (OrderProducts) TableName() string {
	return "order_products"
}
