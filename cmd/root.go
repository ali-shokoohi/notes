package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// RootCmd is the root command!
var RootCmd = &cobra.Command{
	Use:   "",
	Short: "default command",
	Long:  `no functionality`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to SolName!")
	},
}

// Execute runs the main command of the project
func Execute() {
	cobra.OnInitialize(func() {
		// I you have some initial functions you can call them here!
	})
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetConfigPath(cmd *cobra.Command) string {
	configFlag := cmd.Flags().Lookup("config")
	if configFlag != nil {
		configFilePath := configFlag.Value.String()
		log.Println("config file path", zap.String("configFilePath", configFilePath))
		if configFilePath != "" {
			return configFilePath
		}
	}

	return ""
}

func init() {
	RootCmd.PersistentFlags().String("config", "", "path to config file")
}
