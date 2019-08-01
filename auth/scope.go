package auth

import (
	"net/http"
	"strings"

	"github.com/Golang_Northwind_API/utils"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func HasPermission(scope string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
			token := authHeaderParts[1]
			hasScope := checkScope(scope, token)

			if !hasScope {
				message := "Insufficient permissions"
				utils.ResponseJSON(message, w, http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)

		}
		return http.HandlerFunc(fn)
	}
}

func checkScope(scope string, tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, nil)
	claims, _ := token.Claims.(*CustomClaims)
	hasScope := false

	for _, permission := range claims.Permissions {
		if permission == scope {
			hasScope = true
		}
	}

	return hasScope
}
