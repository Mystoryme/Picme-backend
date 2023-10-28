package main

import (
	"picme-backend/modules/config"
	"picme-backend/modules/fiber"
	"picme-backend/modules/minio"
	"picme-backend/modules/mysql"
)

// fiber run สุดท้ายนะ
func main() {
	config.Init()
	mysql.Init()
	minio.Init()
	fiber.Init()

}
