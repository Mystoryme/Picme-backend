package mod

import (
	"gorm.io/gorm"
	"picme-backend/types/module"
)

var Conf *module.Config
var DB *gorm.DB
