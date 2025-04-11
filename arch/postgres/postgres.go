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

func Connect() (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("❌ Failed to connect to database:", err)
		return nil, err
	}

	log.Println("✅ Connected to PostgreSQL successfully")
	return db, nil
}

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

func (db *database) Connect() {}

func (db *database) Disconnect() {}
