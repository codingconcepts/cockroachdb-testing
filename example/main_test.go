package example

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbTests *bool
	db      *sql.DB
)

func TestMain(m *testing.M) {
	dbTests = flag.Bool("db", false, "run database integration tests")
	flag.Parse()

	fmt.Println("run database tests", *dbTests)

	if *dbTests {
		setupDatabase()
	}

	exitCode := m.Run()

	if *dbTests {
		teardownDatabase()
	}

	os.Exit(exitCode)
}

func setupDatabase() {
	connStr, ok := os.LookupEnv("CONN_STR")
	if !ok {
		log.Fatal("connection string env var not found")
	}

	var err error
	if db, err = sql.Open("postgres", connStr); err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error pinging the database: %v", err)
	}
}

func teardownDatabase() {
	db.Close()
}
