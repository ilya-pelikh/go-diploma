package identity

import (
	"errors"
	"os"
)

const JWT_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbiI6InByaXZldCJ9.37gqWJV0QtJL6ll1oiBUZzDb3T-oBvSV7JWx83coozA"

func CheckAuth(dto AuthRequestDTO) (*AuthResponseDTO, error) {
	password := os.Getenv("TODO_PASSWORD")

	if dto.Password == password {
		resp := AuthResponseDTO{
			Token: JWT_TOKEN,
		}
		return &resp, nil
	}
	return nil, errors.New("Пароли не совпадают")
}
