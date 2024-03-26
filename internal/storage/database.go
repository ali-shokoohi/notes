package storage

import (
	"github.com/ali-shokoohi/notes/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	GormDB *gorm.DB
}

func NewGormDB(cfg *config.Config) (*Database, error) {
	var database Database
	dsn := cfg.Database.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		cfg.Logger.Sugar().Errorf("failed to connect to database: %v", err)
		return &database, err
	}
	database.GormDB = db
	return &database, nil
}
