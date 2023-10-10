package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/endpoints/account"
	"picme-backend/endpoints/post"
	profileEndpoint "picme-backend/endpoints/profile"
	"picme-backend/modules/fiber/middlewares"
)

func Init(router fiber.Router) {
	account := router.Group("account/")
	account.Post("register", accountEndpoint.RegisterHandler)
	account.Post("login", accountEndpoint.LoginHandler)

	profile := router.Group("profile/", middlewares.Jwt())
	profile.Get("info", profileEndpoint.ProfileGetHandler)

	post := router.Group("post/")
	post.Get("/list", postEndpoint.GetHandler)
	post.Post("/create", postEndpoint.CreateHandler)
}
