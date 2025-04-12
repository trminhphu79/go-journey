package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	GetInstance() *database
	Connect()
	Disconnect()
}

type PostgresConfig struct {
	User        string
	Pwd         string
	Host        string
	Port        string
	Name        string
	MinPoolSize uint16
	MaxPoolSize uint16
	Timeout     time.Duration
}

type database struct {
	*gorm.DB
	config  PostgresConfig
	context context.Context
}

func CreateDatabase(ctx context.Context, config PostgresConfig) Database {
	db := database{
		context: ctx,
		config:  config,
	}
	return &db
}

func (db *database) GetInstance() *database {
	return db
}

func (db *database) Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Sprintln("host: ", host)
	fmt.Sprintln("user: ", user)
	fmt.Sprintln("port: ", port)
	fmt.Sprintln("password: ", password)
	fmt.Sprintln("dbname: ", dbname)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port,
	)

	var gormDB *gorm.DB
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)

		// Wait before retrying (exponential backoff)
		retryDelay := time.Duration(1<<uint(i)) * time.Second
		if retryDelay > 10*time.Second {
			retryDelay = 10 * time.Second
		}

		log.Printf("Retrying in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	if err != nil {
		log.Println("❌ Maximum retries reached. Failed to connect to database:", err)
		return
	}

	db.DB = gormDB
	log.Println("✅ Connected to PostgreSQL successfully")
}

func (db *database) Disconnect() {}
