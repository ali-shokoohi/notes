// internal/config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Port            int             `mapstructure:"port"`
	SSL             bool            `mapstructure:"ssl"`
	MaintenanceMode bool            `mapstructure:"maintenanceMode"`
	Debug           bool            `mapstructure:"debug"`
	ReleaseMode     string          `mapstructure:"releaseMode"`
	Database        *DatabaseConfig `mapstructure:"database"`
	Logger          *zap.Logger
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Initialize the logger
	if configuration.Debug {
		configuration.Logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Unable to load zap development logger, %v", err)
		}
	} else {
		configuration.Logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Unable to load zap production logger, %v", err)
		}
	}

	return &configuration, nil
}
