// write some code that will;
// connect to the database
// retrieve the first row in the database
// log the retrieved data
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5430
	user     = "postgres"
	password = "dev"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Info().Msg("Unable to connect to db")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Insert new row
	sqlStatement := `
  INSERT INTO stock (id, amount_available, amount_sold, description, location, name, price)
  VALUES ('1251b2c9-6779-4246-b8ab-beb7057d3bc1', 100, 15, 'Xps 13 9310', 'Nigeria', 'Laptop', 1800.0)`
	value, err := db.Exec(sqlStatement)
	log.Info().Msgf("Inserted %d row(s)", value)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
