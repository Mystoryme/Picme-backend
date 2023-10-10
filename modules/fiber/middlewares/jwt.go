package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
	"picme-backend/types/misc"
	"picme-backend/types/response"
)

func Jwt() fiber.Handler {
	conf := jwtware.Config{
		SigningKey:  []byte("babycomeandtakemylovenadruinit"),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ContextKey:  "l",
		Claims:      new(misc.UserClaim),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.Error(false, "JWT validation failure", err)
		},
	}

	return jwtware.New(conf)
}
