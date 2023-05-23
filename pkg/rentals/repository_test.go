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

	t.Run("searchRentals returns the correct rentals when using ID search.", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{Ids: "1,2"}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		if len(rentals) != 2 {
			t.Fatalf("failed to get the proper rentals. expected:2, got: %d", len(rentals))
		}
	})

	t.Run("searchRentals returns the correct rentals when using price_min search.", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{PriceMin: 20000}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		if len(rentals) != 4 {
			t.Fatalf("failed to get the proper rentals. expected:4, got: %d", len(rentals))
		}
	})

	t.Run("searchRentals returns the correct rentals when using price_max search.", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{PriceMax: 10000}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		if len(rentals) != 9 {
			t.Fatalf("failed to get the proper rentals. expected:9, got: %d", len(rentals))
		}
	})

	t.Run("searchRentals returns the correct rentals when using price_min and price_max search.", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{PriceMin: 10000, PriceMax: 15000}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		if len(rentals) != 10 {
			t.Fatalf("failed to get the proper rentals. expected:10, got: %d", len(rentals))
		}
	})

	t.Run("searchRentals returns the correct order of rentals when using sort.", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{Sort: "name"}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		expected := "1984 Volkswagen Westfalia"
		if rentals[0].Name != expected {
			t.Fatalf("failed to get the proper rentals. expected:%s, got: %s", expected, rentals[0].Name)
		}
	})

	t.Run("searchRentals returns the correct order of rentals when using sort is desc", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{Sort: "name", Order: "desc"}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		expected := "TiKi Van  Extended custom camper"
		if rentals[0].Name != expected {
			t.Fatalf("failed to get the proper rentals. expected:%s, got: %s", expected, rentals[0].Name)
		}
	})

	t.Run("searchRentals returns the correct rentals when using near", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{Near: "33.64,-117.93"}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		expected := 6
		if len(rentals) != expected {
			t.Fatalf("failed to get the proper rentals. expected:%d, got: %d", expected, len(rentals))
		}
	})

	t.Run("searchRentals returns the correct number when limited", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{Limit: 7}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		expected := 7
		if len(rentals) != expected {
			t.Fatalf("failed to get the proper rentals. expected:%d, got: %d", expected, len(rentals))
		}
	})

	t.Run("searchRentals returns the correct rental when offset", func(t *testing.T) {
		rentalRepo := setup(t)
		defer tearDown(rentalRepo)

		params := RentalSearchParams{Order: "name", Offset: 5}
		rentals, err := rentalRepo.searchRentals(&params)

		if err != nil {
			t.Fatalf("error searching rentals by ids: %s", err)
		}

		expected := "'Abaco' VW Bay Window: Westfalia Pop-top"
		if rentals[0].Name != expected {
			t.Fatalf("failed to get the proper rentals. expected:%s, got: %s", expected, rentals[0].Name)
		}
	})
}
