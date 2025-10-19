package infra

import (
	domainUser "gohabits/internal/domain/user"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDatabase membuat koneksi baru ke PostgreSQL menggunakan GORM
func NewDatabase(cfg *Config) *gorm.DB {
	if cfg.DBDsn == "" {
		log.Fatal("❌ DB_DSN belum diatur, pastikan environment variable DB_DSN ada")
	}

	db, err := gorm.Open(postgres.Open(cfg.DBDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Gagal konek ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Gagal ambil instance sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//auto migrate
	err = db.AutoMigrate(&domainUser.User{})
	if err != nil {
		log.Fatalf("❌ Gagal melakukan auto migrate: %v", err)
	}

	log.Println("✅ Database connected successfully")
	return db
}
