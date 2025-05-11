package config

// DB Конфиг для БД
type Mongo struct {
	Host     string `envconfig:"DB_HOST,default=mindlab-mongodb"`
	Port     string `envconfig:"DB_PORT,default=27017"`
	Username string `envconfig:"DB_USERNAME,default=default"`
	Password string `envconfig:"DB_PASSWORD,default=default"`
	DBName   string `envconfig:"DB_DBNAME,default=default"`
	SSLMode  string `envconfig:"DB_SSLMODE,default=default"`
}
