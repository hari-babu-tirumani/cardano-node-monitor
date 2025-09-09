package config

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	NodeURL  string        `mapstructure:"node_url"`
	Port     string        `mapstructure:"port"`
	Interval time.Duration `mapstructure:"interval"`
	LogLevel string        `mapstructure:"log_level"`
}

func Load(cmd *cobra.Command) (*Config, error) {
	v := viper.New()
	
	v.SetDefault("node_url", "http://localhost:12798")
	v.SetDefault("port", "8080")
	v.SetDefault("interval", 30*time.Second)
	v.SetDefault("log_level", "info")

	configFile, _ := cmd.Flags().GetString("config")
	if configFile != "" {
		v.SetConfigFile(configFile)
	} else {
		v.SetConfigName(".cardano-monitor")
		v.SetConfigType("yaml")
		v.AddConfigPath("$HOME")
		v.AddConfigPath(".")
	}

	v.AutomaticEnv()
	v.SetEnvPrefix("CARDANO_MONITOR")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	nodeURL, _ := cmd.Flags().GetString("node-url")
	if nodeURL != "" {
		v.Set("node_url", nodeURL)
	}

	port, _ := cmd.Flags().GetString("port")
	if port != "" {
		v.Set("port", port)
	}

	interval, _ := cmd.Flags().GetDuration("interval")
	if interval > 0 {
		v.Set("interval", interval)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}