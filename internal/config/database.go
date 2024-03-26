package config

import (
	"fmt"
)

type DatabaseConfig struct {
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

func (dc *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dc.Host, dc.User, dc.Password, dc.Database, dc.Port)
}
