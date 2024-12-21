package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5433 user=zhora password=300 dbname=shelfshare_test sslmode=disable"
	}

	os.Exit(m.Run())
}
