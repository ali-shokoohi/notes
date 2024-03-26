package cmd

import (
	"log"

	"github.com/ali-shokoohi/notes/internal/config"
	"github.com/ali-shokoohi/notes/internal/model"
	"github.com/ali-shokoohi/notes/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate ",
	Long:  `Migrate `,
	RunE:  migrateCmdE,
}

func migrateCmdE(cmd *cobra.Command, args []string) error {
	configPath := GetConfigPath(cmd)
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("failed at loading config err:", zap.Error(err))
	}
	cfg.Logger.Info("setting up migrate command")

	// Loading the database
	db, err := storage.NewGormDB(cfg)
	if err != nil {
		cfg.Logger.Error("failed at creating a new Gorm database:", zap.Error(err))
		return err
	}

	return db.GormDB.AutoMigrate(model.Note{})
}

func init() {
	RootCmd.AddCommand(migrateCMD)

	//postgres flags
	migrateCMD.Flags().String("postgres_note", "amir", "Define postgres note")
	viper.BindPFlag("database.note", migrateCMD.Flags().Lookup("postgres_note"))

	migrateCMD.Flags().String("postgres_pwd", "amir123", "Define postgres db password")
	viper.BindPFlag("database.password", migrateCMD.Flags().Lookup("postgres_pwd"))

	migrateCMD.Flags().String("postgres_db", "datalead", "Define postgres db name")
	viper.BindPFlag("database.database", migrateCMD.Flags().Lookup("postgres_db"))

	migrateCMD.Flags().String("postgres_host", "localhost", "Define postgres host address .e.g localhost")
	viper.BindPFlag("database.host", migrateCMD.Flags().Lookup("postgres_host"))

	migrateCMD.Flags().Int("postgres_port", 5432, "Define postgres host address .e.g localhost")
	viper.BindPFlag("database.port", migrateCMD.Flags().Lookup("postgres_port"))
}
