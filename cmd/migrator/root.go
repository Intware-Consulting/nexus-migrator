package migrator

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "nexus-migrator",
		Short: "nexus-migrator - simple CLI tool to migrate repositories inside a single Nexus instance",
		Long: `nexus-migrator is a simple CLI tool created to migrate repositories inside a single Nexus instance.
	
You can use nexus-migrator to migrate, import and export repositories contents`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func Execute() {
	cobra.OnInitialize(initConfig)

	initLogger()

	rootCmd.AddCommand(migrateCmd, downloadCmd, uploadCmd)
	rootCmd.Execute()
}

func initConfig() {
	viper.SetDefault("level", "info")
	viper.SetConfigFile("nexus.conf")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Msg("Using config file: " + viper.ConfigFileUsed())
	} else {
		log.Info().Msg("Using env variables configuration")
		viper.AutomaticEnv()

	}
	zerolog.SetGlobalLevel(getLevel(viper.GetString("level")))
}

func initLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339})
}

func getLevel(level string) zerolog.Level {
	mapping := map[string]zerolog.Level{
		"trace":    zerolog.TraceLevel,
		"debug":    zerolog.DebugLevel,
		"info":     zerolog.InfoLevel,
		"warn":     zerolog.WarnLevel,
		"error":    zerolog.ErrorLevel,
		"fatal":    zerolog.FatalLevel,
		"panic":    zerolog.PanicLevel,
		"disabled": zerolog.Disabled,
	}

	return mapping[strings.ToLower(level)]
}
