package dto

type UserLogin struct {
	ID           uint   `json:"id"`
	NamaPengguna string `json:"nama_pengguna"`
	Username     string `json:"username"`
	Nik          string `json:"nik"`
	Role         string `json:"role"`
}

type LoginResponse struct {
	User  UserLogin `json:"user"`
	Token string    `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
