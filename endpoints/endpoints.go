package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/endpoints/account"
	commentEndpoint "picme-backend/endpoints/comment"
	"picme-backend/endpoints/post"
	profileEndpoint "picme-backend/endpoints/profile"
	"picme-backend/modules/fiber/middlewares"
)

func Init(router fiber.Router) {
	account := router.Group("account/")
	account.Post("register", accountEndpoint.RegisterHandler)
	account.Post("login", accountEndpoint.LoginHandler)

	profile := router.Group("profile/", middlewares.Jwt())
	profile.Get("/info", profileEndpoint.ProfileGetHandler)
	profile.Get("/post", profileEndpoint.ProfilePostGetHandler)

	post := router.Group("post/", middlewares.Jwt())
	post.Get("/list", postEndpoint.GetHandler)
	post.Post("/create", postEndpoint.CreateHandler)
	post.Post("/like", postEndpoint.LikeHandler)

	comment := router.Group("comment/", middlewares.Jwt())
	comment.Post("/create", commentEndpoint.CreateHandler)
	comment.Get("/list", commentEndpoint.GetHandler)

}
