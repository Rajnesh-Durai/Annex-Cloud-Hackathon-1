package config

import (
	"time"

	"github.com/spf13/viper"
)
//mapstructure - It's a library that helps in mapping arbitrary data into Go structures
//				 specify how environment variables should be mapped to struct fields
type Config struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`

	ServerPort string `mapstructure:"PORT"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
// When the application starts, it calls the LoadConfig function to load the configuration.
// The configuration can be stored in environment variables or a configuration file.
// Viper reads the configuration and maps it to the Config struct using mapstructure.
// The application can then use the Config struct to access configuration parameters throughout its execution.