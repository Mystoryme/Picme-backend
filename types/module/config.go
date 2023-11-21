package module

type Config struct {
	Address        string   `yaml:"address" validate:"required"`
	Cors           []string `yaml:"cors" validate:"required"`
	MysqlDsn       string   `yaml:"mysql_dsn" validate:"required"`
	BucketEndpoint string   `yaml:"bucket_endpoint" validate:"required"`
	BucketName     string   `yaml:"bucket_name" validate:"required"`
	BucketKey      string   `yaml:"bucket_key"  validate:"required"`
	BucketSecret   string   `yaml:"bucket_secret" validate:"required"`
	AutoMigrate    bool     `yaml:"auto_migrate"`
	ScbUrl         string   `yaml:"scb_url" validate:"required"`
	ScbAppKey      string   `yaml:"scb_app_key" validate:"required"`
	ScbAppSecret   string   `yaml:"scb_app_secret" validate:"required"`
}
