package orders_entity

import (
	"time"
)

type ResponseOrderDetail struct {
	ID                 int        `gorm:"primaryKey" json:"id"`
	RefCode            string     `json:"ref_code"`
	UserName           string     `json:"user_name"`
	AddressName        string     `json:"address_name"`
	AddressNo          string     `json:"address_no"`
	AddressMoo         string     `json:"address_moo"`
	AddressRoad        string     `json:"address_road"`
	AddressSubDistrict string     `json:"address_sub_district"`
	AddressDistrict    string     `json:"address_district"`
	AddressProvince    string     `json:"address_province"`
	AddressZipCode     string     `json:"address_zip_code"`
	AddressTel         string     `json:"address_tel"`
	CompletedDate      *time.Time `json:"completed_date"`
	ShippedDate        *time.Time `json:"shipped_date"`
	CancelDate         *time.Time `json:"cancel_date"`
	OrderStatus        string     `json:"order_status"`
	SummaryPrice       int        `json:"summary_price"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	UpdatedBy          string     `json:"updated_by"`
}
