package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/endpoints/account"
	"picme-backend/endpoints/post"
)

func Init(app fiber.Router) {
	account := app.Group("account/")
	account.Post("/login", accountEndpoint.LoginHandler)

	post := app.Group("post/")
	post.Get("/list", postEndpoint.GetHandler)
	post.Post("/create", postEndpoint.CreateHandler)
}
