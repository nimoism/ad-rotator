package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	intapp "github.com/nimoism/ad-rotator/internal/app"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ad-rotator",
		Short: "Start ad-rotator server",
		RunE:  serve,
	}
	configFile string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.toml", "Path to config file")
	rootCmd.AddCommand(migrateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func serve(_ *cobra.Command, _ []string) error {
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := initConfig(configFile)
	if err != nil {
		return err
	}

	app, err := intapp.NewApp(config)
	if err != nil {
		return err
	}
	return app.Run(ctx)
}

func initConfig(filename string) (intapp.Config, error) {
	var config intapp.Config
	viper.SetEnvPrefix("adr")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("config reading error: %w", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("config unmarshaling error: %w", err)
	}
	return config, nil
}
