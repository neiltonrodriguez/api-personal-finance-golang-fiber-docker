package auth

import (
	"context"
	"errors"
	"personal-finance-api/domain"
	LoginModel "personal-finance-api/models/login"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUserNotFound = errors.New("Usuário não encontrado")
)

func SignIn(ctx context.Context, email, password string) (string, domain.User, error) {
	user, err := LoginModel.GetUserByEmail(ctx, email)
	if err != nil {
		return "", domain.User{}, err
	}
	if user.Id == 0 {
		return "", domain.User{}, ErrUserNotFound
	}
	err = domain.VerifyPassword(user.Password, password)
	if err != nil {
		return "", domain.User{}, err
	}

	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(domain.GlobalConfig.JwtSecretKey))
	if err != nil {
		return "", domain.User{}, err
	}
	user.Password = ""
	return t, user, nil
}

func CheckToken(ctx context.Context) (bool, error) {

	// idm := ctx.Value("token")
	// user, err := UserModel.GetById()
	return true, nil
}
