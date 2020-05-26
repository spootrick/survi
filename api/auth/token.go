package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/response"
	"github.com/spootrick/survi/config"
	"net/http"
	"time"
)

// GenerateJWT generates new token for client
func GenerateJWT(user model.User) (string, error) {
	claims := model.JWTClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Survi App",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SecretKey)
}

// ExtractJWT extracts the token from headers
func ExtractJWT(w http.ResponseWriter, r *http.Request) *jwt.Token {
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&model.JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return config.SecretKey, nil
		},
	)

	if err != nil {
		code := http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				err = errors.New("your token has expired")
				response.ERROR(w, code, err)
				return nil
			case jwt.ValidationErrorSignatureInvalid:
				err = errors.New("the signature is invalid")
				response.ERROR(w, code, err)
				return nil
			default:
				response.ERROR(w, code, err)
				return nil
			}
		default:
			err = errors.New("authentication token required")
			response.ERROR(w, http.StatusUnauthorized, err)
		}
	}

	return token
}
