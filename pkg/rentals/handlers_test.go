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

		expected := `{"id":1,"name":"'Abaco' VW Bay Window: Westfalia Pop-top","description":"ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi","type":"camper-van","make":"Volkswagen","model":"Bay Window","year":"1978","length":15,"sleeps":"4","Primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg","price":{"day":16900},"location":{"city":"Costa Mesa","state":"CA","zip":"92627","country":"US","lat":33.64,"lng":-117.93},"user":{"id":1,"first_name":"John","last_name":"Smith"}}`
		if recorder.Body.String() != expected {
			t.Errorf("handler returned unexpected body. expected: %v \ngot: %v",
				expected, recorder.Body.String())
		}
	})

	t.Run("extractParams correctly marshals url params into struct", func(t *testing.T) {

		request, err := http.NewRequest("GET", "/rentals?ids=1,2&limit=10&offset=5&near=34,32&price_min=0&price_max=100&sort=name&order=desc", nil)
		request.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		params, err := extractParams(request)
		if err != nil {
			t.Fatalf("failed to extract params: %s", err)
		}
		expectedName := "1,2"
		if params.Ids != expectedName {
			t.Errorf("handler returned unexpected body. expected: %s got: %s",
				expectedName, params.Ids)
		}
		expectedMin := 0
		if params.PriceMin != expectedMin {
			t.Errorf("handler returned unexpected price_min. expected: %d got: %d",
				expectedMin, params.PriceMin)
		}
		expectedMax := 100
		if params.PriceMax != expectedMax {
			t.Errorf("handler returned unexpected price_max. expected: %d got: %d",
				expectedMax, params.PriceMax)
		}
		expectedNear := "34,32"
		if params.Near != expectedNear {
			t.Errorf("handler returned unexpected near. expected: %s got: %s",
				expectedNear, params.Near)
		}
		expectedSort := "name"
		if params.Sort != expectedSort {
			t.Errorf("handler returned unexpected near. expected: %s got: %s",
				expectedSort, params.Sort)
		}
		expectedOrder := "desc"
		if params.Order != expectedOrder {
			t.Errorf("handler returned unexpected order. expected: %s got: %s",
				expectedOrder, params.Order)
		}
		expectedLimit := 5
		if params.Limit != expectedLimit {
			t.Errorf("handler returned unexpected limit. expected: %d got: %d",
				expectedLimit, params.Limit)
		}
		expectedOffset := 0
		if params.Offset != expectedOffset {
			t.Errorf("handler returned unexpected offset. expected: %d got: %d",
				expectedOffset, params.Offset)
		}
	})

	t.Run("test SearchRentals", func(t *testing.T) {
		rentalService := setupHandlers(t)
		defer teardownHandlers(rentalService)
		request, err := http.NewRequest("GET", "/rentals?price_min=24000", nil)
		request.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(rentalService.SearchRentalsJSON)

		handler.ServeHTTP(recorder, request)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `[{"id":10,"name":"Betty!    1987 Volkswagen Westfalia Poptop Manual with kitchen!","description":"mollis curabitur cum convallis sagittis feugiat lectus ligula porta libero parturient maecenas cum facilisis ridiculus mauris ut est scelerisque tincidunt quisque hac lectus mus dapibus","type":"camper-van","make":"Volkswagen","model":"Westfalia","year":"1987","length":15,"sleeps":"4","Primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1535836865/p/rentals/91133/images/blijuwlisflua72ay1p2.jpg","price":{"day":25000},"location":{"city":"Missoula ","state":"MT","zip":"59808","country":"US","lat":46.92,"lng":-114.09},"user":{"id":5,"first_name":"Ben","last_name":"Reynard"}}]`
		if recorder.Body.String() != expected {
			t.Errorf("handler returned unexpected body. expected: %v \ngot: %v",
				expected, recorder.Body.String())
		}
	})
}
