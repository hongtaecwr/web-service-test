package database_entity

import "time"

type Addresses struct {
	ID                 int        `gorm:"primaryKey" json:"id"`
	Name               string     `json:"name"`
	AddressNo          string     `json:"address_no"`
	AddressMoo         string     `json:"address_moo"`
	AddressRoad        string     `json:"address_road"`
	AddressSubDistrict string     `json:"address_sub_district"`
	AddressDistrict    string     `json:"address_district"`
	AddressProvince    string     `json:"address_province"`
	AddressZipCode     string     `json:"address_zip_code"`
	AddressTel         string     `json:"address_tel"`
	UserId             int        `json:"user_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `sql:"index" json:"deleted_at"`
	CreatedBy          string     `json:"created_by"`
	UpdatedBy          string     `json:"updated_by"`
	DeletedBy          string     `json:"deleted_by"`
}

func (Addresses) TableName() string {
	return "addresses"
}
