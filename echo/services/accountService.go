package services

import (
	"app/models"
	"app/repositories"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type accountService struct {
	accountRepository repositories.AccountRepository
}

func NewAccountService(ar repositories.AccountRepository) *accountService {
	return &accountService{accountRepository: ar}
}

// Secrets Manager KeyVaultなどからSDKで取得する
var jwtSecretKey = []byte("jwt_secret_key")

func (as accountService) CreateUser(user *models.UserAuthRequest) (int64, error) {
	user.Password = hashPassword(user.Password)
	return as.accountRepository.InsertUser(user)
}

func (as accountService) GetAllUsernames() ([]string, error) {

	return as.accountRepository.SelectAllUsernames()
}

func (as accountService) Verify(c echo.Context) bool {
	cookie, err := c.Cookie("token")
	if err != nil {
		return false
	}

	tokenString := cookie.Value

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return false
	}

	return true
}

func (as accountService) CreateJWT(userID int64) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Id:        strconv.Itoa(int(userID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as accountService) Login(request *models.UserAuthRequest) (int64, error) {
	request.Password = hashPassword(request.Password)
	user := as.accountRepository.GetUser(request.Username)
	if user == nil {
		return 0, errors.New("user not found")
	}
	if user.Password != request.Password {
		return 0, errors.New("incorrect password")
	}
	return user.ID, nil
}

func (as accountService) CreateBeingDeletedCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.MaxAge = -1

	return cookie
}

func (as accountService) DeleteUser(userID int64) error {
	return as.accountRepository.Delete(userID)
}

// accountServiceのみでしか使用しないためプライベート
func hashPassword(pw string) string {
	hash := sha256.New()
	hash.Write([]byte(pw))
	hashedPassword := hash.Sum(nil)

	return hex.EncodeToString(hashedPassword)
}
