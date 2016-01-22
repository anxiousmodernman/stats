package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	csvfile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Could not open csv file for write.")
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	writer.Comma = '|' // Use this as delimiter, aka the "Comma"

	query := `SELECT first_name, last_name
	FROM people WHERE last_name != "McFarland"`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatal(err)
	}

	// Make a lil index guy
	i := 0
	// loop over rows
	for rows.Next() {
		// Initialize two vars as sql.NullString (if we might get nulls)
		// Define your output here. NULLABLE things have their own type.
		var firstName sql.NullString
		var lastName sql.NullString

		// The & says, "Give me a pointer to this". See p. 32 of Go book.
		err := rows.Scan(&firstName, &lastName)
		if err != nil {
			log.Fatalln(err)
		}
		// the csvwriter Write() method requires a []string so we make one.
		toWrite := []string{firstName.String, lastName.String}
		writer.Write(toWrite)

		// Periodically flush to disk every 1000 iterations
		if i%1000 == 0 {
			writer.Flush()
		}
	}
	// Flush any remaining data to disk.
	writer.Flush()
}
