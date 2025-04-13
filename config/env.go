package config

import (
	"os"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	// Server
	GoMode     string `mapstructure:"GO_MODE"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort uint16 `mapstructure:"SERVER_PORT"`

	// Database
	DBHost         string `mapstructure:"DB_HOST"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBMinPoolSize  uint16 `mapstructure:"DB_MIN_POOL_SIZE"`
	DBMaxPoolSize  uint16 `mapstructure:"DB_MAX_POOL_SIZE"`
	DBQueryTimeout uint16 `mapstructure:"DB_QUERY_TIMEOUT"`

	// keys
	RSAPublicKey  string `mapstructure:"RSA_PUBLIC_KEY"`  // This matches your docker-compose.yml
	RSAPrivateKey string `mapstructure:"RSA_PRIVATE_KEY"` // This matches your docker-compose.yml

	// Token
	AccessTokenValiditySec  uint64 `mapstructure:"ACCESS_TOKEN_VALIDITY_SEC"`
	RefreshTokenValiditySec uint64 `mapstructure:"REFRESH_TOKEN_VALIDITY_SEC"`
}

var (
	envInstance *Env
	envOnce     sync.Once
)

// GetEnv implements lazy loading of environment variables
func GetEnv() *Env {
	envOnce.Do(func() {
		envInstance = loadEnv(".env", true)
	})
	return envInstance
}

// loadEnv loads environment variables from file or environment
func loadEnv(filename string, override bool) *Env {
	env := Env{}

	// Set default values
	viper.SetDefault("GO_MODE", "development")
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", 8080)
	viper.SetDefault("DB_MIN_POOL_SIZE", 10)
	viper.SetDefault("DB_MAX_POOL_SIZE", 100)
	viper.SetDefault("DB_QUERY_TIMEOUT", 30)
	viper.SetDefault("ACCESS_TOKEN_VALIDITY_SEC", 3600)
	viper.SetDefault("REFRESH_TOKEN_VALIDITY_SEC", 604800)

	log.WithFields(log.Fields{
		"DB_HOST_ENV": os.Getenv("DB_HOST"),
		"DB_PORT_ENV": os.Getenv("DB_PORT"),
		"DB_NAME_ENV": os.Getenv("DB_NAME"),
		"DB_USER_ENV": os.Getenv("DB_USER"),
	}).Info("Raw environment variables")

	// Inside loadEnv function, after viper.AutomaticEnv() but before unmarshal
	// Explicitly bind environment variables
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("RSA_PUBLIC_KEY")
	viper.BindEnv("RSA_PRIVATE_KEY")

	// Enable reading from environment variables
	viper.AutomaticEnv()

	// Try to load from .env file if it exists
	if _, err := os.Stat(filename); err == nil {
		viper.SetConfigFile(filename)
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Warn("Could not read config file, using environment variables")
		} else {
			log.Info("Loaded configuration from file: ", filename)
		}
	} else {
		log.Info(".env file not found — using environment variables")
	}

	// Unmarshal configuration into struct
	if err := viper.Unmarshal(&env); err != nil {
		log.WithError(err).Fatal("Failed to parse environment configuration")
	}

	// Legacy support for APP_PORT and GIN_MODE
	if os.Getenv("APP_PORT") != "" && env.ServerPort == 0 {
		if port, err := strconv.ParseUint(os.Getenv("APP_PORT"), 10, 16); err == nil {
			env.ServerPort = uint16(port)
			log.WithField("port", env.ServerPort).Info("Using APP_PORT for server port")
		}
	}

	if os.Getenv("GIN_MODE") != "" && env.GoMode == "" {
		env.GoMode = os.Getenv("GIN_MODE")
		log.WithField("mode", env.GoMode).Info("Using GIN_MODE for Go mode")
	}

	// Log loaded configuration (excluding sensitive data)
	log.WithFields(log.Fields{
		"GO_MODE":                    env.GoMode,
		"SERVER_HOST":                env.ServerHost,
		"SERVER_PORT":                env.ServerPort,
		"DB_HOST":                    env.DBHost,
		"DB_NAME":                    env.DBName,
		"DB_PORT":                    env.DBPort,
		"DB_USER":                    env.DBUser,
		"DB_PASSWORD":                "********", // Masked for security
		"DB_MIN_POOL_SIZE":           env.DBMinPoolSize,
		"DB_MAX_POOL_SIZE":           env.DBMaxPoolSize,
		"DB_QUERY_TIMEOUT":           env.DBQueryTimeout,
		"RSA_PUBLIC_KEY":             env.RSAPublicKey,
		"RSA_PRIVATE_KEY":            env.RSAPrivateKey,
		"ACCESS_TOKEN_VALIDITY_SEC":  env.AccessTokenValiditySec,
		"REFRESH_TOKEN_VALIDITY_SEC": env.RefreshTokenValiditySec,
	}).Info("Environment configuration loaded")

	return &env
}

// Legacy function for backward compatibility
func IniitEnv(filename string, override bool) *Env {
	return loadEnv(filename, override)
}

func (e *Env) GetDBConnectionString() string {
	return "host=" + e.DBHost +
		" user=" + e.DBUser +
		" password=" + e.DBPassword +
		" dbname=" + e.DBName +
		" port=" + e.DBPort +
		" sslmode=disable TimeZone=UTC"
}
