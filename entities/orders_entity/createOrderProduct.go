package orders_entity

import "time"

type CreateOrderProduct struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	OrderId   int        `json:"order_id"`
	ProductId int        `json:"product_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
}

func (CreateOrderProduct) TableName() string {
	return "order_products"
}
