package infra

import (
	"gorm.io/gorm"
)

// Container menampung dependency global (Config, DB, dll)
type Container struct {
	Config    *Config
	DB        *gorm.DB
	Redis     *RedisClient
	Validator *Validator
}
type Validator struct{}

// NewContainer membuat container untuk dependency injection
func NewContainer() *Container {
	cfg := LoadConfig()
	// init validator
	InitValidator()

	return &Container{
		Config:    cfg,
		DB:        NewDatabase(cfg),
		Redis:     NewRedisClient(cfg),
		Validator: &Validator{},
	}
}
