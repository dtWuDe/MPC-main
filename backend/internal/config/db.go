package config

type DBConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"tinh"`
	Password string `env:"DB_PASSWORD" envDefault:"123"`
	DBName   string `env:"DB_NAME" envDefault:"mpc_key"`
}
