package mod

import (
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"picme-backend/types/module"
)

var Minio *minio.Client
var Conf *module.Config
var DB *gorm.DB
