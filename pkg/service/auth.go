package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	salt   = "2312312313"
	secret = "asdafdsggdqwerdsfffasdasxsd"
)

type AuthService struct {
	repo repository.Authentication
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func NewAuthService(repository *repository.Repository) *AuthService {
	return &AuthService{
		repo: repository.Authentication,
	}
}

func (s *AuthService) VerifyCode(ctx context.Context, code string) (string, error) {
	output, err := s.repo.VerifyCode(ctx, code)
	if err != nil {
		return "", err
	}

	userID, err := s.repo.CreateUser(ctx, models.User{Phone: output.Phone})
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(15 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	generatedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return generatedToken, err
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not type of *tokenClaims")
	}

	return claims.UserID, nil
}

func generateSixDigitNumber() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func sendMessage(phone string, code string) {

	body := models.SMSRequestBody{
		APIKey:    os.Getenv("nexmo_api_key"),
		APISecret: os.Getenv("nexmo_api_secret"),
		To:        phone,
		From:      "Region.app",
		Text:      code,
	}

	smsBody, _ := json.Marshal(body)

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
}

func (s *AuthService) SendSMS(ctx context.Context, phone string) (string, error) {
	generatedCode := generateSixDigitNumber()

	//go sendMessage(phone, generatedCode)
	//if err != nil {
	//	log.Println(err.Error())
	//	return "", err
	//}

	code := models.Code{
		Phone: phone,
		Code: generatedCode,
		ExpiresAt: time.Now().Add(120*time.Second),
	}

	err := s.repo.CreateCode(ctx, code)
	if err != nil{
		log.Println(err.Error())
		return "", err
	}

	return generatedCode, nil
}



func (s *AuthService) SignIn(ctx context.Context, user models.User) (string, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *AuthService) SignUp(ctx context.Context, user models.User) (string, error) {
	return s.repo.GetUser(ctx, user)
}

func generatePassword(password string)  string {
	return ""
}