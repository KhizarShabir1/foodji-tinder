package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func InitDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",

		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	//creating sql tables
	CreateTables(db)

	return db, nil
}

// CreateTables create all three sql tables for product, session, and vote
func CreateTables(db *sql.DB) error {

	stmt := `
        CREATE TABLE IF NOT EXISTS product (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        );

        CREATE TABLE IF NOT EXISTS session (
            id TEXT PRIMARY KEY,
            created_at TIMESTAMP NOT NULL
        );

        CREATE TABLE IF NOT EXISTS vote (
            id SERIAL PRIMARY KEY,
            session_id TEXT NOT NULL,
            product_id INTEGER NOT NULL,
            liked BOOLEAN NOT NULL,
            FOREIGN KEY (session_id) REFERENCES session (id),
            FOREIGN KEY (product_id) REFERENCES product (id)
        );
    `
	_, err := db.Exec(stmt)
	if err != nil {
		log.Println("Error creating tables")
		panic(err)
	}

	return nil
}
