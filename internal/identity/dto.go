package identity

type AuthRequestDTO struct {
	Password string `json:"password" validate:"required"`
}

type AuthResponseDTO struct {
	Token string `json:"token"`
}
