package product_entity

type ResponseProducts struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Sold        int    `json:"sold"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

func (ResponseProducts) TableName() string {
	return "products"
}
