package rentals

import (
	"database/sql"
	"testing"
)

func setupService(t *testing.T) *RentalService {

	db, err := sql.Open("postgres", "postgres://root:root@172.18.0.2/testingwithrentals?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to DB: %s", err)
	}

	rentalRepo := NewRentalRepository(db)

	return NewRentalService(rentalRepo)
}

func teardownService(service *RentalService) {
	service.repository.database.Close()
}

func TestUserService(t *testing.T) {

	t.Run("GetRental returns a rental", func(t *testing.T) {
		rentalService := setupService(t)
		defer teardownService(rentalService)
		rental, err := rentalService.GetRental(1)
		if err != nil {
			t.Fatalf("error finding rental: %s", err)
		}
		expectedName := "'Abaco' VW Bay Window: Westfalia Pop-top"
		if rental.Name != expectedName {
			t.Fatalf("did not return a correct rental, expected: %s, got:%s", expectedName, rental.Name)
		}
	})

	t.Run("SearchRentals returns a rental", func(t *testing.T) {
		rentalService := setupService(t)
		defer teardownService(rentalService)
		params := &RentalSearchParams{Ids: "2"}
		rental, err := rentalService.SearchRentals(params)
		if err != nil {
			t.Fatalf("error finding rental: %s", err)
		}
		expectedName := "Maupin: Vanagon Camper"
		if rental[0].Name != expectedName {
			t.Fatalf("did not return a correct rental, expected: %s, got:%s", expectedName, rental[0].Name)
		}
	})

}
