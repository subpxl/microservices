package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func SetupDB() *pgx.Conn {
	// connString := "host=postgres user=todoservice dbname=todoservice password=todoservice sslmode=disable"
	connString := os.Getenv("POSTGRES_URI")
	if connString == "" {
		log.Fatal("cannot find psql db connection string")
	}
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Println("cannot connect to db , ", err)
	}

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("PostgreSQL version: %s\n", version)

	return conn
}

func WaitForDB() {
	maxAttempts := 10
	for i := 1; i <= maxAttempts; i++ {
		err := DB.Ping(context.Background())
		if err == nil {
			fmt.Println("Database is ready!")
			return
		}

		fmt.Printf("Attempt %d: Unable to ping database. Retrying...\n", i)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Failed to connect to the database after multiple attempts.")
}
func SyncDB() {
	_, err := DB.Exec(context.Background(), `create table if not exists Todo(
		id SERIAL PRIMARY KEY,
		Todo varchar(255)
	);`)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("cretetd successfully")
}
