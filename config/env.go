package config

import (
	"log"
	"os"
	"strconv"

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
	RSAPublicKeyPath  string `mapstructure:"RSA_PUBLIC_KEY_PATH"`
	RSAPrivateKeyPath string `mapstructure:"RSA_PRIVATE_KEY_PATH"`

	// Token
	AccessTokenValiditySec  uint64 `mapstructure:"ACCESS_TOKEN_VALIDITY_SEC"`
	RefreshTokenValiditySec uint64 `mapstructure:"REFRESH_TOKEN_VALIDITY_SEC"`
}

func IniitEnv(filename string, override bool) *Env {
	env := Env{}
	viper.SetConfigFile(filename)

	viper.SetDefault("GO_MODE", "development")
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", 8080)
	viper.SetDefault("DB_MIN_POOL_SIZE", 10)
	viper.SetDefault("DB_MAX_POOL_SIZE", 100)
	viper.SetDefault("DB_QUERY_TIMEOUT", 30)
	viper.SetDefault("ACCESS_TOKEN_VALIDITY_SEC", 3600)
	viper.SetDefault("REFRESH_TOKEN_VALIDITY_SEC", 604800)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Warning: Error reading env file:", err)
		log.Println("Using environment variables only")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error loading environment config:", err)
	}

	if os.Getenv("APP_PORT") != "" && env.ServerPort == 0 {
		if port, err := strconv.ParseUint(os.Getenv("APP_PORT"), 10, 16); err == nil {
			env.ServerPort = uint16(port)
		}
	}

	if os.Getenv("GIN_MODE") != "" && env.GoMode == "" {
		env.GoMode = os.Getenv("GIN_MODE")
	}

	return &env
}

func (e *Env) GetDBConnectionString() string {
	return "host=" + e.DBHost +
		" user=" + e.DBUser +
		" password=" + e.DBPassword +
		" dbname=" + e.DBName +
		" port=" + e.DBPort +
		" sslmode=disable TimeZone=UTC"
}
