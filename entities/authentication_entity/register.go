package authentication

import "time"

type RequestRegister struct {
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	TelNo      string    `json:"tel_no"`
	UserStatus string    `json:"user_status"`
	Permission int       `json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
}

func (RequestRegister) TableName() string {
	return "users"
}
