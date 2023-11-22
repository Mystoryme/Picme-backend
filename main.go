package main

import (
	"picme-backend/helper"
	"picme-backend/modules/config"
)

// fiber run สุดท้ายนะ
func main() {
	config.Init()
	helper.ScbCreateQrPayment()
	//mysql.Init()
	//minio.Init()
	//fiber.Init()

}
