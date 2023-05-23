package rentals

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func setup(t *testing.T) *rentalRepository {
	db, err := sql.Open("postgres", "postgres://root:root@172.18.0.2/testingwithrentals?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to DB: %s", err)
	}

	return NewRentalRepository(db)
}

func tearDown(rentalRepo *rentalRepository) {
	rentalRepo.database.Close()
}

func TestRepository(t *testing.T) {

	t.Run("getRental returns the correct rental.", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		rental := rentalRepo.getRental(1)
		expectedName := "'Abaco' VW Bay Window: Westfalia Pop-top"
		if rental.Name != expectedName {
			t.Fatalf("failed to get the rental. expected:%s, got: %s", expectedName, rental.Name)
		}
	})
}
