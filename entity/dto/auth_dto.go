package dto

type AuthRequestDto struct {
	User     string `json:"email"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
}
