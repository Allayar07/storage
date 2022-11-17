package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"storage/internal/model"
	"storage/internal/repository"
	"time"
)

const (
	salt    = "y6gt397623iusndhc8jkbldbj89jkjs"
	signKey = "rtcnvdvjb83745blkjzvkjbfdkvjbvf"
)

type AuthService struct {
	repos repository.Authorization
}

type TokenClaims struct {
	jwt.StandardClaims
	Userid int `json:"userid"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repos: repo,
	}
}
func (s *AuthService) Create(u model.User) (int, error) {
	u.Password = HashPassword(u.Password)
	return s.repos.Create(u)
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repos.GetUser(username, HashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signKey))
}
func (s *AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not that type")
	}
	return claims.Userid, nil
}
func HashPassword(password string) string {
	hash := sha1.New()

	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
