package user_entity

type ResponseUserDetail struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	TelNo      string `json:"tel_no"`
	UserStatus string `json:"user_status"`
	Permission int    `json:"permission"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

func (ResponseUserDetail) TableName() string {
	return "users"
}
