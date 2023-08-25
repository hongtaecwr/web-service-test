package database_entity

import "time"

type Users struct {
	ID         int        `gorm:"primaryKey" json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	TelNo      string     `json:"tel_no"`
	UserStatus string     `json:"user_status"`
	Permission int        `json:"permission"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	CreatedBy  string     `json:"created_by"`
	UpdatedBy  string     `json:"updated_by"`
	DeletedBy  string     `json:"deleted_by"`
}

func (Users) TableName() string {
	return "users"
}

const (
	ActiveStatus    = "ACTIVE"
	InactiveStatus  = "INACTIVE"
	PermissionUser  = 1
	PermissionAdmin = 2
)
