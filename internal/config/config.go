package config

import (
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	AWSRegion          string
	SQSQueueURL        string
	DBConnectionString string
}

// LoadConfig loads the configuration from environment variables or a configuration file
func LoadConfig() (*Config, error) {
	viper.SetEnvPrefix("your_module")
	viper.AutomaticEnv()

	// Set default values if necessary
	viper.SetDefault("AWS_REGION", "us-west-2")

	// Bind environment variables
	viper.BindEnv("AWS_REGION")
	viper.BindEnv("SQS_QUEUE_URL")
	viper.BindEnv("DB_CONNECTION_STRING")

	// Read values from environment variables or a configuration file
	// For example, you can use viper.ReadInConfig() to read from a file

	// Create the Config object
	cfg := &Config{
		AWSRegion:          viper.GetString("AWS_REGION"),
		SQSQueueURL:        viper.GetString("SQS_QUEUE_URL"),
		DBConnectionString: viper.GetString("DB_CONNECTION_STRING"),
	}

	return cfg, nil
}
