package mysql

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	mod "picme-backend/modules"
	"picme-backend/types/table"
	"time"
)

func Init() {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open SQL connection
	dialector := mysql.New(
		mysql.Config{
			DSN:               mod.Conf.MysqlDsn,
			DefaultStringSize: 255,
		},
	)

	// * Open main database connection
	if db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	}); err != nil {
		logrus.Fatal("UNABLE TO LOAD GORM MYSQL DATABASE")
	} else {
		mod.DB = db
	}
	migrate()
	logrus.Debug("INITIALIZED MYSQL CONNECTION")
}

func migrate() error {
	// * Migrate table
	if err := mod.DB.AutoMigrate(
		new(table.PostComments),
		new(table.PostLikes),
		new(table.PostViews),
		new(table.Posts),
		new(table.Users),
		new(table.PostBookMark),
		new(table.PostBoost),
		new(table.PostDonate),
		new(table.Message)); err != nil {
		return err

	}

	return nil
}
