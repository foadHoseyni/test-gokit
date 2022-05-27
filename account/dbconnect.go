package account

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// postgresql function to return DB connection
func GetDBconn() *sql.DB {

	// const dbsource = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db
}
