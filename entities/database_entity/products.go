package database_entity

import "time"

type Products struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	Price       int        `json:"price"`
	Description string     `json:"description"`
	Stock       int        `json:"stock"`
	Sold        int        `json:"sold"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedBy   string     `json:"updated_by"`
	DeletedBy   string     `json:"deleted_by"`
}
