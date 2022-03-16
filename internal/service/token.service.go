package service

import (
	"authentication/internal/apperrors"
	"authentication/internal/models"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	jwtKey = []byte(os.Getenv("JWT_KEY"))
)

type TokenService interface {
	GenerateJWT(user *models.User) (string, error)
	ParseJWT(token string) (*models.User, error)
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (t *tokenService) GenerateJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["firstName"] = user.FirstName
	claims["lastName"] = user.LastName
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", apperrors.GeneratingTokenErr
	}

	return tokenString, nil
}

func (t *tokenService) ParseJWT(userToken string) (*models.User, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, apperrors.NotValidTokenErr
	}

	var currentUser *models.User
	currentUser = &models.User{
		Email: claims["email"].(string),
	}

	// handle cases where values are empty
	for k, v := range claims {
		switch k {
		case "firstName":
			if v == nil {
				currentUser.FirstName = ""
			} else {
				currentUser.FirstName = v.(string)
			}
		case "lastName":
			if v == nil {
				currentUser.LastName = ""
			} else {
				currentUser.LastName = v.(string)
			}
		case "username":
			if v == nil {
				currentUser.Username = ""
			} else {
				currentUser.Username = v.(string)
			}
		}

	}

	return currentUser, nil
}
