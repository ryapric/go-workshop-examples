// I don't have a good resource for this one, except the stdlib docs
// (specifically the `encoding/json` & `database/sql` packages)
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Underscores here are for "side-effect" imports -- this package provides a
	// SQLite driver, and that's it
	_ "github.com/mattn/go-sqlite3"
)

// When reading in external data (like JSON, or DB rows), you will most often
// see a struct defining its schema & types -- which can either be strict like
// this, or loose with a map[string]interface{}/any. If you *can* be strict, do
// it. Struct tags are optionally how you can map incoming fields to fields in
// the Go struct itself -- they're super handy!

// SalesRecord defines a single record in the incoming sales data JSON array
type SalesRecord struct {
	Date        string   `json:"date"`
	Region      string   `json:"region"`
	CustomerID  string   `json:"customer_id"`
	Revenue     float64  `json:"revenue"`
	AddlDetails []string `json:"addl_details,omitempty"` // `omitempty` means if it's not in the record, Go will ignore it, and accessing it will return `nil`
}

// AllSalesData represents the array of sales records themselves, which is how
// they come in. json.Unmarshal() will actually handle a full parse of the
// array, though, so this is just here as a reference convenience
type AllSalesRecords []SalesRecord

func getSalesRecordsFromJSONBytes(rawData []byte) AllSalesRecords {
	var records AllSalesRecords
	err := json.Unmarshal(rawData, &records)
	if err != nil {
		log.Fatalf("error unmarshaling raw JSON data: %s\n", err.Error())
	}
	return records
}

func writeRecordsToDB(records AllSalesRecords) {
	// opens up the DB connection
	db, err := sql.Open("sqlite3", "./sales.db")
	if err != nil {
		log.Fatalf("error opening database: %s\n", err.Error())
	}
	defer db.Close()

	// Applies the DDL to match the SalesRecord struct
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sales (
			date TEXT,
			region TEXT,
			customer_id TEXT,
			revenue NUMBER
		);
	`)
	if err != nil {
		log.Fatalf("error executing DDL: %s\n", err.Error())
	}

	// Prepares a statement that we'll pass values to -- almost all languages
	// have facilities for this, and you should use them
	stmt, err := db.Prepare(`INSERT INTO sales VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Fatalf("error preparing base SQL statement: %s\n", err.Error())
	}

	// Loop over all the individual records, and write them out to the table
	for _, record := range records {
		_, err = stmt.Exec(record.Date, record.Region, record.CustomerID, record.Revenue)
		if err != nil {
			log.Fatalf("error inserting rows into database: %s\n", err.Error())
		}
	}
}

func main() {
	// NOTE: this DB removal is just here so you get clean runs for this POC.
	// IRL, you'll want to do something else that makes sense
	os.Remove("./sales.db")

	// First script arg (e.g. go run main.go <arg1>) is the JSON file containing
	// sales data
	if len(os.Args) < 2 {
		log.Fatalln("you must provide a sales data file path as a CLI arg")
	}
	rawDataFile := os.Args[1]

	// Read in the data from disk
	rawData, err := os.ReadFile(rawDataFile)
	if err != nil {
		log.Fatalf("error when reading file: %s\n", err.Error())
	}

	// Call the two ETL funcs we defined above (well, here it's just E & L, but)
	records := getSalesRecordsFromJSONBytes(rawData)
	writeRecordsToDB(records)

	log.Println("Successfully wrote JSON data to SQLite database! If you have the SQLite3 CLI installed, you can check the contents yourself by running the following:")
	log.Println("sqlite3 -header -column ./sales.db 'select * from sales;'")
}
