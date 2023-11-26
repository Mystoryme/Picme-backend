package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"picme-backend/endpoints"
	mod "picme-backend/modules"
	"time"
)

func Init() {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  ErrorHandler,
		AppName:       "Picme",
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	app.Use(recover.New())

	// Initialize API endpoints
	apiGroup := app.Group("api/")
	endpoints.Init(apiGroup)

	// * Run the server
	err := app.Listen(mod.Conf.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
