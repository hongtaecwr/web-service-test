package database_entity

import "time"

type OrderLogs struct {
	ID            int        `gorm:"primaryKey" json:"id"`
	RefCode       string     `json:"ref_code"`
	UserId        int        `json:"user_id"`
	AddressId     int        `json:"address_id"`
	CompletedDate *time.Time `json:"completed_date"`
	ShippedDate   *time.Time `json:"shipped_date"`
	CancelDate    *time.Time `json:"cancel_date"`
	OrderStatusId int        `json:"order_status_id"`
	SummaryPrice  int        `json:"summary_price"`
	OrderId       int        `json:"order_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at"`
	CreatedBy     string     `json:"created_by"`
	UpdatedBy     string     `json:"updated_by"`
	DeletedBy     string     `json:"deleted_by"`
}

func (OrderLogs) TableName() string {
	return "order_logs"
}
