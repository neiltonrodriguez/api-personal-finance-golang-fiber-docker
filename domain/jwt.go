package domain

// import (
// 	"html"
// 	"net/http"
// 	"strings"
// 	"time"

// 	jwt "github.com/dgrijalva/jwt-go"
// )

// var secretKey = GlobalConfig.JwtSecretKey

// func GenerateJWT(user User) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["user_email"] = user.Email
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
// 	return token.SignedString(secretKey)
// }

// func JwtExtract(r *http.Request) (map[string]interface{}, error) {
// 	tokenString := ExtractBearerToken(r)
// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return claims, nil
// }

// func ExtractBearerToken(r *http.Request) string {
// 	headerAuthorization := r.Header.Get("Authorization")
// 	bearerToken := strings.Split(headerAuthorization, " ")
// 	return html.EscapeString(bearerToken[1])
// }
