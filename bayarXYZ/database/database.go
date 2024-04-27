package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewEnvDBConfig() *DBConfig {
	return &DBConfig{
		host:     "localhost",
		port:     "3306",
		username: "root",
		password: "8imaaaTyo09",
		database: "todoapp_development",
	}
}

func (c *DBConfig) GetHost() string {
	return c.host
}

func (c *DBConfig) GetPort() string {
	return c.port
}

func (c *DBConfig) GetUsername() string {
	return c.username
}

func (c *DBConfig) GetPassword() string {
	return c.password
}

func (c *DBConfig) GetDatabase() string {
	return c.database
}

func ConnectToDB() (*sql.DB, error) {
	// Get a database handle.
	config := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "bayarxyz_development",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db, err
}
