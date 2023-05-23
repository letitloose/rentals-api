package rentals

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupHandlers(t *testing.T) *RentalService {

	db, err := sql.Open("postgres", "postgres://root:root@172.18.0.2/testingwithrentals?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to DB: %s", err)
	}

	rentalRepo := NewRentalRepository(db)

	return NewRentalService(rentalRepo)
}

func teardownHandlers(service *RentalService) {
	service.repository.database.Close()
}

func TestHandlers(t *testing.T) {

	t.Run("test GetRental", func(t *testing.T) {
		rentalService := setupHandlers(t)
		defer teardownHandlers(rentalService)
		request, err := http.NewRequest("GET", "/rentals/1", nil)
		request.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(rentalService.GetRentalJSON)

		handler.ServeHTTP(recorder, request)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"id":1,"name":"'Abaco' VW Bay Window: Westfalia Pop-top","description":"ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi","type":"camper-van","make":"Volkswagen","model":"Bay Window","year":"1978","length":15,"sleeps":"4","Primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg","price":{"day":16900},"location":{"city":"Costa Mesa","state":"CA","zip":"92627","country":"US","lat":33.64,"lng":-117.93},"user":{"id":1,"first_name":0,"last_name":0}}`
		if recorder.Body.String() != expected {
			t.Errorf("handler returned unexpected body. expected: %v \ngot: %v",
				expected, recorder.Body.String())
		}
	})
}
