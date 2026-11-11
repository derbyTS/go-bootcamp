package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"restapi/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type ContextString string

func JWTMiddleware(next http.Handler) http.Handler {
	fmt.Println("------------- JWT Middleware -------------")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("------------- Inside JWT Middleware -------------")
		cookie, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return
		}
		token := cookie.Value
		fmt.Println("token:", token)
		jwtSecret := os.Getenv("JWT_SECRET")

		// This is sample from https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(jwtSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			// log.Fatal(err) //This is default

			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Valid {
			log.Println("Valid JWT")
		} else {
			http.Error(w, "Invalid JWT Token", http.StatusUnauthorized)
			log.Println("Invalid JWT: ", cookie.Value)
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if ok {
			fmt.Println(claims["uid"], claims["exp"], claims["role"])
		} else {
			http.Error(w, "Invalid JWT Token", http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), ContextString("role"), claims["role"])
		// These 3 claims argument is from pkg/utils/jwt.go
		ctx = context.WithValue(ctx, ContextString("expiresAt"), claims["exp"])
		ctx = context.WithValue(ctx, ContextString("username"), claims["user"])
		ctx = context.WithValue(ctx, ContextString("userID"), claims["uid"])

		fmt.Println(ctx)

		next.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println("------------- Sent response from JWT Middleware -------------")
	})
}
