package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/luozi-csu/lzblogs/model"
	"github.com/pkg/errors"
)

const (
	Issuer = "lzblogs"
)

type UserClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type JWTService struct {
	SigningKey     []byte
	Issuer         string
	ExpireDuration time.Duration
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		SigningKey:     []byte(secret),
		Issuer:         Issuer,
		ExpireDuration: 24 * time.Hour,
	}
}

func (s *JWTService) NewToken(user *model.User) (string, error) {
	claims := UserClaims{
		ID:   user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(s.ExpireDuration),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.Issuer,
			NotBefore: time.Now().Unix() - int64(10*time.Second),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(s.SigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s *JWTService) ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return s.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &model.User{
		ID:   claims.ID,
		Name: claims.Name,
	}, nil
}
