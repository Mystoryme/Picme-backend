package endpoints

import (
	accountEndpoint "picme-backend/endpoints/account"
	commentEndpoint "picme-backend/endpoints/comment"
	insightEndpoint "picme-backend/endpoints/insight"
	notificationEndpoint "picme-backend/endpoints/notification"
	postEndpoint "picme-backend/endpoints/post"
	profileEndpoint "picme-backend/endpoints/profile"
	"picme-backend/modules/fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	account := router.Group("account/")
	account.Post("register", accountEndpoint.RegisterHandler)
	account.Post("login", accountEndpoint.LoginHandler)

	profile := router.Group("profile/", middlewares.Jwt())
	profile.Get("/info", profileEndpoint.ProfileGetHandler)
	profile.Get("/post", profileEndpoint.ProfilePostGetHandler)
	profile.Get("/gridpost", profileEndpoint.GridPostGetHandler)
	profile.Get("/bookmarkpost", profileEndpoint.BookmarkPostGetHandler)
	profile.Post("/donate", profileEndpoint.DonateHandler)
	profile.Post("/donate/inquiry", profileEndpoint.DonateUserInquiry)
	profile.Post("/edit", profileEndpoint.EditHandler)
	profile.Post("/search", profileEndpoint.SearchHandler)
	profile.Post("/profile_search", profileEndpoint.ProfileSearchGetHandler)
	profile.Post("/post_search", profileEndpoint.PostSearchGetHandler)
	profile.Post("/grid_search", profileEndpoint.GridSearchGetHandler)
	profile.Post("/avatar", profileEndpoint.AvatarHandler)

	post := router.Group("post/", middlewares.Jwt())
	post.Get("/list", postEndpoint.GetHandler)
	post.Post("/create", postEndpoint.CreateHandler)
	post.Post("/like", postEndpoint.LikeHandler)
	post.Delete("/unlike", postEndpoint.UnLikeHandler)
	post.Delete("/unbookmark", postEndpoint.UnBookmarkHandler)
	post.Post("/bookmark", postEndpoint.BookmarkHandler)
	post.Delete("/delete", postEndpoint.DeleteHandler)
	post.Post("/donate", postEndpoint.DonateHandler)
	post.Post("/donate/inquiry", postEndpoint.DonateInquiry)
	post.Post("/view", postEndpoint.ViewHandler)
	post.Post("/boost", postEndpoint.BoostHandler)
	post.Post("/boost/inquiry", postEndpoint.BoostInquiry)
	post.Get("/listboost", postEndpoint.GetBoostHandler)
	post.Post("/post_get", postEndpoint.GetPostHandler)

	comment := router.Group("comment/", middlewares.Jwt())
	comment.Post("/create", commentEndpoint.CreateHandler)
	comment.Post("/list", commentEndpoint.GetHandler)
	comment.Delete("/delete", commentEndpoint.DeleteHandler)

	notification := router.Group("notification/", middlewares.Jwt())
	notification.Get("/list", notificationEndpoint.GetHandler)

	insight := router.Group("insight/", middlewares.Jwt())
	insight.Get("/", insightEndpoint.GetHandler)
	insight.Post("/search", insightEndpoint.GetSearch)
	insight.Post("/view_profile", insightEndpoint.ViewProfile)
	insight.Get("/search_history", insightEndpoint.GetViewHistory)
}
