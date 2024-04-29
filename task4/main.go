package main

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v4"
)

const (
	connString = "postgres://postgres:mysecretpassword@localhost:5432/postgres"
)

func main() {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS test_table (
			id SERIAL PRIMARY KEY,
			value TEXT
		)
	`); err != nil {
		fmt.Printf("Unable to create table: %v\n", err)
	}

	insertRows(conn, 10000)
	insertRows(conn, 100000)
	insertRows(conn, 1000000)

	queryRows(conn)
}

func insertRows(conn *pgx.Conn, count int) {
	start := time.Now()
	fmt.Printf("Inserting %d rows...\n", count)
	for i := 0; i < count; i++ {
		value := fmt.Sprintf("value_%d", i)
		if _, err := conn.Exec(context.Background(), "INSERT INTO test_table (value) VALUES ($1)", value); err != nil {
			fmt.Printf("Error inserting rows: %v\n", err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Inserted %d rows in %v\n", count, elapsed)
}

func queryRows(conn *pgx.Conn) {
	start := time.Now()
	fmt.Printf("Querying rows...")
	rows, err := conn.Query(context.Background(), "SELECT * FROM test_table")
	if err != nil {
		fmt.Printf("Error querying rows: %v\n", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error reading rows: %v\n", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Queried %d rows in %v\n", count, elapsed)
}
