package module

type Config struct {
	Address  string   `yaml:"address" validate:"required"`
	Cors     []string `yaml:"cors" validate:"required"`
	MysqlDsn string   `yaml:"mysql_dsn" validate:"required"`
}
