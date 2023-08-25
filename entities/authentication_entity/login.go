package authentication

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
