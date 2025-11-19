package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *Config, loglrus *logrus.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.New(
		// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
		// 	logger.Config{
		// 		SlowThreshold: time.Second, // Query di atas 1 detik dianggap lambat
		// 		LogLevel:      logger.Info, // Bisa diganti ke Warn atau Error jika terlalu verbose
		// 		Colorful:      true,
		// 	},
		// ),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err = connection.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected and pinged PostgreSQL database successfully")

	// Konfigurasi koneksi pool
	maxIdleConns := 30
	if cfg.SetMaxIdleConns != "" {
		maxIdleConns, _ = strconv.Atoi(cfg.SetMaxIdleConns)
	}

	maxOpenConns := 100
	if cfg.SetMaxOpenConns != "" {
		maxOpenConns, _ = strconv.Atoi(cfg.SetMaxOpenConns)
	}

	maxLifeTimeConnection := 300
	if cfg.SetMaxLifeTime != "" {
		maxLifeTimeConnection, _ = strconv.Atoi(cfg.SetMaxLifeTime)
	}

	connection.SetMaxIdleConns(maxIdleConns)
	connection.SetMaxOpenConns(maxOpenConns)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	if os.Getenv("ENV") == "development" {
		db = db.Debug()
	}

	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close connection to PostgreSQL:", err)
	}
	dbSQL.Close()
}
