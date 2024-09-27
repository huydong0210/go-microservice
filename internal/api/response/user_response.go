package response

type UserInfoResponse struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	HashPassword string `json:"hash_password"`
	Roles        string `json:"roles"`
	Email        string `json:"email"`
}
