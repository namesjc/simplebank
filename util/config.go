package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the apaplication.
// The values are read by viper from a config file or environment variables.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml, yaml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
