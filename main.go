package main

import (
	"picme-backend/modules/config"
	"picme-backend/modules/fiber"
	"picme-backend/modules/mysql"
)

func main() {
	config.Init()
	mysql.Init()
	fiber.Init()
}
