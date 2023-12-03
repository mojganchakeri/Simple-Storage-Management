package models

type Env struct {
	ServerAddress          string `mapstructure:"SERVER_ADDRESS" yaml:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT" yaml:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST" yaml:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT" yaml:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER" yaml:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS" yaml:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME" yaml:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR" yaml:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR" yaml:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET" yaml:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET" yaml:"REFRESH_TOKEN_SECRET"`
	LogLevel               string `mapstructure:"LOG_LEVEL" yaml:"LOG_LEVEL"`
	StoreServiceHost       string `mapstructure:"STORE_SERVICE_HOST" yaml:"STORE_SERVICE_HOST"`
	StoreServicePort       string `mapstructure:"STORE_SERVICE_PORT" yaml:"STORE_SERVICE_PORT"`
}
