package infra

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Redis struct {
	Addr     string
	Password string
	DB       int
}

type Config struct {
	AppName   string
	Port      string
	JWTSecret string
	DBDsn     string
	ENV       string
	Redis     Redis
}

// Singleton instance agar konfigurasi hanya di-load sekali
var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		// Implementation for loading configuration
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		cfg = &Config{
			ENV:       GetEnv("APP_ENV", "development"),
			AppName:   GetEnv("APP_NAME", "GoHabits"),
			Port:      GetEnv("PORT", "3000"),
			JWTSecret: GetEnv("JWT_SECRET", "supersecretkey"),
			DBDsn:     GetEnv("DB_DSN", ""),
			Redis: Redis{
				Addr:     GetEnv("REDIS_ADDR", "localhost:6379"),
				Password: GetEnv("REDIS_PASSWORD", ""),
				DB:       0, // default DB
			},
		}
	})

	return cfg
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() *Config {
	if cfg == nil {
		return LoadConfig()
	}
	return cfg
}
