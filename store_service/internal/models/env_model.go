package models

type Env struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	LogLevel       string `mapstructure:"LOG_Level"`
	StoragePath    string `mapstructure:"STORAGE_PATH"`
	MaxStorageSize int    `mapstructure:"MAX_STORAGE_SIZE"`
	MaxFileSize    int    `mapstructure:"MAX_FILE_SIZE"`
}
