package models

type Env struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS" yaml:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT" yaml:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST" yaml:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT" yaml:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER" yaml:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS" yaml:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME" yaml:"DB_NAME"`
	LogLevel       string `mapstructure:"LOG_Level" yaml:"LOG_Level"`
	StoragePath    string `mapstructure:"STORAGE_PATH" yaml:"STORAGE_PATH"`
	MaxStorageSize int    `mapstructure:"MAX_STORAGE_SIZE" yaml:"MAX_STORAGE_SIZE"`
	MaxFileSize    int    `mapstructure:"MAX_FILE_SIZE" yaml:"MAX_FILE_SIZE"`
}
