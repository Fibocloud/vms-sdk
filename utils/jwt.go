package utils

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// access secret key
func GetAccessSecretKey() []byte {
	return []byte(viper.GetString("APP_SECRET"))
}

type TokenClaims struct {
	ID uint `json:"id"`
	jwt.MapClaims
}

// ExtractJWTString Get claim from token string
func ExtractJWTString(tokenString string) (TokenClaims, error) {
	retClaim := TokenClaims{}
	JwtToken, err := jwt.ParseWithClaims(tokenString, retClaim, func(t *jwt.Token) (interface{}, error) {
		return GetAccessSecretKey(), nil
	})
	if err == nil {
		if !JwtToken.Valid {
			return retClaim, nil
		}
	}
	return retClaim, err
}

// GenerateToken ...
func GenerateToken(userId uint) string {
	accessExpTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": accessExpTime.Unix(),
	}

	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(GetAccessSecretKey())
	return accessToken
}

// GenerateHash password hash generate
func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword compare password and hash
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomWithCharset(length int, charset []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomUpperCase(l int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	return RandomWithCharset(l, letters)
}

func RandomString(l int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	return RandomWithCharset(l, letters)
}

func RandomNumber(l int) string {
	var letters = []rune("0123456789")
	return RandomWithCharset(l, letters)
}
