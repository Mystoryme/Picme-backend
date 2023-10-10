package crypto

import (
	"github.com/golang-jwt/jwt/v4"
	"picme-backend/types/response"
)

func SignJwt(claim jwt.Claims) (string, *response.ErrorInstance) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	if signedToken, err := token.SignedString([]byte("babycomeandtakemylovenadruinit")); err != nil {
		return "", response.Error(true, "Unable to sign JWT token", err)
	} else {
		return signedToken, nil
	}
}
