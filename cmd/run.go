package cmd

import (
	"log"

	"github.com/ali-shokoohi/notes/http"
	"github.com/ali-shokoohi/notes/internal/config"
	"github.com/ali-shokoohi/notes/internal/handler"
	gormRepository "github.com/ali-shokoohi/notes/internal/repository/gorm"
	"github.com/ali-shokoohi/notes/internal/route"
	"github.com/ali-shokoohi/notes/internal/service"
	"github.com/ali-shokoohi/notes/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run ",
	Long:  `Run `,
	RunE:  runCmdE,
}

func runCmdE(cmd *cobra.Command, args []string) error {
	configPath := GetConfigPath(cmd)
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("failed at loading config err:", zap.Error(err))
	}
	cfg.Logger.Info("setting up run command")

	// Log to file in production mode
	// if !cfg.Debug {
	// 	if err := initProdLog(gormConfig, &logConfig{
	// 		gorm: true,
	// 		gin:  true,
	// 		zap:  true,
	// 	}); err != nil {
	// 		return err
	// 	}
	// }

	cfg.Logger.Info("setup rbac middleware")
	// rbacMiddleware, err := auth.NewRBACMiddleware("configs/auth_model.conf", "configs/routes_policy.csv", cfg.Auth0.Namespace) //FIXME
	// if err != nil {
	// 	cfg.Logger.Error("unable to setup rbac middleware", zap.String("error", err.Error()))
	// 	return errors.Newf("unable to setup rbac middleware: %v", err)
	// }

	if cfg.Debug {
		cfg.Logger.Sugar().Infof("setting gin debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		cfg.Logger.Sugar().Infof("setting gin release mode")
		gin.SetMode(gin.ReleaseMode)
	}

	// Loading the database
	db, err := storage.NewGormDB(cfg)
	if err != nil {
		cfg.Logger.Error("failed at creating a new Gorm database:", zap.Error(err))
		return err
	}

	// Define repositories
	noteRepository := gormRepository.NewNoteRepository(db.GormDB)

	// Define services
	noteService := service.NewNoteService(cfg, noteRepository)

	// Define handlers
	generalHandler := handler.NewGeneralHandler()
	noteHandler := handler.NewNoteHandler(cfg, noteService)

	// Define gin.Engine and routes
	r := gin.New()
	route.LoadRoutes(r, generalHandler, noteHandler)
	// server := server.NewServer(authenticator, rbacMiddleware, db,
	// 	&cfg.OAuthConfig, &server.ServerConfig{})
	server := http.NewServer(cfg, r)

	return server.Launch()

}

func init() {
	RootCmd.AddCommand(runCMD)

	// port
	runCMD.Flags().Uint("port", 5050, "HTTP server listen address")
	viper.BindPFlag("port", runCMD.Flags().Lookup("port"))

	//auth keys
	// TODO: Remove if not needed
	runCMD.Flags().String("auth_pkey", "", "jwt authentication private key")
	runCMD.Flags().String("auth_pubkey", "", "jwt authentication public key")

	//postgres flags
	runCMD.Flags().String("postgres_note", "amir", "Define postgres note")
	viper.BindPFlag("database.note", runCMD.Flags().Lookup("postgres_note"))

	// TODO: Remove if not needed
	runCMD.Flags().String("postgres_connect_name", "", "Define postgres connect name from gcp")

	runCMD.Flags().String("postgres_pwd", "amir123", "Define postgres db password")
	viper.BindPFlag("database.password", runCMD.Flags().Lookup("postgres_pwd"))

	runCMD.Flags().String("postgres_db", "datalead", "Define postgres db name")
	viper.BindPFlag("database.database", runCMD.Flags().Lookup("postgres_db"))

	runCMD.Flags().String("postgres_host", "localhost", "Define postgres host address .e.g localhost")
	viper.BindPFlag("database.host", runCMD.Flags().Lookup("postgres_host"))

	runCMD.Flags().Int("postgres_port", 5432, "Define postgres host address .e.g localhost")
	viper.BindPFlag("database.port", runCMD.Flags().Lookup("postgres_port"))

	// log flags
	runCMD.Flags().String("log_path", "", "Define log path")
	viper.BindPFlag("log_path", runCMD.Flags().Lookup("log_path"))

	// ssl flags
	runCMD.Flags().Bool("ssl", false, "Define ssl")
	viper.BindPFlag("ssl", runCMD.Flags().Lookup("ssl"))

	// debug flags
	runCMD.Flags().Bool("debug", false, "Define debug")

}
