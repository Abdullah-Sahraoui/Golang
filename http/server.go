package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "windows"
	dbname   = "face_recognition_db"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, dbErr := sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		panic(dbErr)
	}

	// Create a new context and begin a transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	// tx above is an instance of `*sql.Tx` through which we can execute our queries

	// Here, the query is executed on the transaction instance, and not applied to the database yet
	_, tErr := tx.ExecContext(ctx, "INSERT INTO Users (name, email, joined) VALUES ('Stones', 'stones@gmail.com', current_date);")
	if tErr != nil {
		tx.Rollback()
		panic(tErr)
		return
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Success!")
}
