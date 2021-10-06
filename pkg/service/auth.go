package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"io"
	"log"
	"net/http"
	"os"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repository *repository.Repository) *AuthService {
	return &AuthService{
		repo: repository.Authentication,
	}
}

func (s *AuthService) VerifyCode(ctx context.Context, code string) error {
	return s.repo.VerifyCode(ctx, code)
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

func (s *AuthService) SendSMS(ctx context.Context, phone string) (string, error) {
	generatedCode := generateSixDigitNumber()

	code := models.Code{
		Phone: phone,
		Code: generatedCode,
	}

	err := s.repo.CreateCode(ctx, code)
	if err != nil{
		log.Println(err.Error())
		return "", err
	}

	body := models.SMSRequestBody{
		APIKey:    os.Getenv("nexmo_api_key"),
		APISecret: os.Getenv("nexmo_api_secret"),
		To:        phone,
		From:      "Region.app",
		Text:      generatedCode,
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	return generatedCode, err
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