package database

type Config struct {
	Name string `env:"DB_NAME" json:", omitempty"`
	User string `env:"DB_USER" json:", omitempty"`
	Password string `env:"DB_PASSWORD json:"-""` // ignored by zap's JSON formatter
	Host string `env:"DB_HOST, default=localhost" json:", omitempty"`
	Port string `env:"DB_PORT, default=5432" json:", omitempty"`
	SSLMode string `env:"DB_SSLMODE, default=require" json:", omitempty"`
}