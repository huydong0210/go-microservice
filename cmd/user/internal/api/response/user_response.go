package response

type UserLoginResponse struct {
	Username     string `json:"username"`
	HashPassword string `json:"hash_password"`
	Roles        string `json:"roles"`
	Email        string `json:"email"`
}
