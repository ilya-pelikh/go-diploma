package identity

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"diploma/internal/pkg/config"
)

func CreateToken(dto AuthRequestDTO) (*AuthResponseDTO, error) {
	if dto.Password != config.ENV.TODO_PASSWORD {
		return nil, errors.New("password is not correct")
	}

	t := jwt.New(jwt.SigningMethodHS256)
	signedToken, err := t.SignedString([]byte(config.ENV.TODO_JWT_KEY))
	if err != nil {
		return nil, errors.New("couldn't create token")
	}

	return &AuthResponseDTO{Token: signedToken}, nil
}

func VerifyToken(tokenString string) (bool, error) {
	parsed, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.ENV.TODO_JWT_KEY), nil
	})
	if err != nil {
		return false, err
	}
	return parsed.Valid, nil
}
