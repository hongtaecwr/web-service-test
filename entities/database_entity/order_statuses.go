package database_entity

type OrderStatuses struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

const (
	IDStatusOrdered   = 1
	IDStatusShipped   = 2
	IDStatusCompleted = 3
	IDStatusCancel    = 4
)
