package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/setting"
)

type Jwt struct {
}

func NewJwt() *Jwt {
	return &Jwt{}
}

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int    `json:"role"`
	jwt.StandardClaims
}

func (obj *Jwt) GenerateToken(id int, name string, email string, role int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(6 * time.Hour)

	claims := Claims{
		id,
		name,
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "fancy-devops",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func (obj *Jwt) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (obj *Jwt) GetClaims(c *gin.Context) (*Claims, error) {
	token := c.GetHeader("token")
	return NewJwt().ParseToken(token)
}
