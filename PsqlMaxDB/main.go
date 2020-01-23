package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/twinj/uuid"
)

var (
	id               int
	connectionString string
	amount           int
)

type Database struct {
	Connection *sql.DB
}

func init() {
	connectionString = os.Getenv("CONN")
	amount, _ = strconv.Atoi(os.Getenv("ROWS"))
}
func main() {
	var db Database

	client, err := InitDB()
	CheckErr(err)
	db.Connection = client
	for i := 0; i < amount; i++ {
		db.InsertRow()

	}
} //end main

func InitDB() (*sql.DB, error) {
	conn, err := sql.Open("postgres", connectionString)
	CheckErr(err)
	if err == nil {
		return conn, nil
	}

	err = errors.New("Failed to connect to database")

	return nil, err
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ConstructQuery(val string) string {
	query := "INSERT INTO test (uuid) values ('" + val + "');"

	return query
}

func (db Database) InsertRow() {
	u := uuid.NewV4()
	query := ConstructQuery(u.String())
	fmt.Printf("\nExecuting: %s", query)
	rows, err := db.Connection.Query(query)
	CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id)
	}
}
