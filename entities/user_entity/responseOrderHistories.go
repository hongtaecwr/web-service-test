package user_entity

import "time"

type ResponseOrderHistories struct {
	ID            int        `gorm:"primaryKey" json:"id"`
	RefCode       string     `json:"ref_code"`
	UserId        int        `json:"user_id"`
	AddressId     int        `json:"address_id"`
	CompletedDate *time.Time `json:"completed_date"`
	ShippedDate   *time.Time `json:"shipped_date"`
	CancelDate    *time.Time `json:"cancel_date"`
	OrderStatusId int        `json:"order_status_id"`
	SummaryPrice  int        `json:"summary_price"`
}

func (ResponseOrderHistories) TableName() string {
	return "orders"
}
