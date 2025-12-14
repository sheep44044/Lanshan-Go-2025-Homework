package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	JWTSecretKey      string        `mapstructure:"JWT_SECRET_KEY"`
	JWTIssuer         string        `mapstructure:"JWT_ISSUER"`
	JWTExpirationTime time.Duration `mapstructure:"JWT_EXPIRATION_TIME"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`
}

func Load() (*Config, error) {
	v := viper.New()
	setDefaults(v)
	configureViper(v)
	if err := readConfiguration(v); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("DB_USER", "root")
	v.SetDefault("DB_PASSWORD", "root")
	v.SetDefault("DB_HOST", "localhost")
	v.SetDefault("DB_PORT", "3306")
	v.SetDefault("DB_NAME", "notes_db")
	v.SetDefault("SERVER_PORT", "8080")
	v.SetDefault("JWT_SECRET_KEY", "lanshan_homework07")
	v.SetDefault("JWT_ISSUER", "note_app")
	v.SetDefault("JWT_EXPIRATION_TIME", "24h")

	v.SetDefault("REDIS_HOST", "localhost")
	v.SetDefault("REDIS_PORT", "6379")
	v.SetDefault("REDIS_PASSWORD", "")
	v.SetDefault("REDIS_DB", "0")
}

func configureViper(v *viper.Viper) {
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func readConfiguration(v *viper.Viper) error {
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("Warning: .env file not found, using defaults and system env")
			return nil
		}
		return fmt.Errorf("config file error: %w", err)
	}
	fmt.Printf("Using config file: %s\n", v.ConfigFileUsed())
	return nil
}
