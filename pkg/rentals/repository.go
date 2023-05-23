package rentals

import (
	"database/sql"
)

type Rental struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	Make            string   `json:"make"`
	Model           string   `json:"model"`
	Year            string   `json:"year"`
	Length          float32  `json:"length"`
	Sleeps          string   `json:"sleeps"`
	PrimaryImageURL string   `json:"Primary_image_url"`
	Price           Price    `json:"price"`
	Location        Location `json:"location"`
	User            User     `json:"user"`
}

type Price struct {
	Day int `json:"day"`
}

type Location struct {
	City      string  `json:"city"`
	State     string  `json:"state"`
	Zip       string  `json:"zip"`
	Country   string  `json:"country"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

type User struct {
	Id        int `json:"id"`
	FirstName int `json:"first_name"`
	LastName  int `json:"last_name"`
}

type rentalRepository struct {
	database *sql.DB
}

func NewRentalRepository(database *sql.DB) *rentalRepository {
	return &rentalRepository{database: database}
}

func (repository *rentalRepository) getRental(id int) *Rental {
	query := "select rentals.id, name, description, type, vehicle_make, vehicle_model, vehicle_year, vehicle_length,  sleeps, primary_image_url, price_per_day, home_city, home_state, home_zip, home_country, lat, lng, u.id, u.first_name, u.last_name from rentals, users u where u.id = user_id and rentals.id = $1"

	rows := repository.database.QueryRow(query, id)

	rental := &Rental{}

	rows.Scan(&rental.Id, &rental.Name, &rental.Description, &rental.Type, &rental.Make, &rental.Model, &rental.Year, &rental.Length, &rental.Sleeps, &rental.PrimaryImageURL,
		&rental.Price.Day, &rental.Location.City, &rental.Location.State, &rental.Location.Zip, &rental.Location.Country, &rental.Location.Latitude, &rental.Location.Longitude,
		&rental.User.Id, &rental.User.FirstName, &rental.User.LastName)

	return rental
}
