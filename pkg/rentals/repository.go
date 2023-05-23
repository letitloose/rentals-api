package rentals

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
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

type rentalSearchParams struct {
	price_min int
	price_max int
	limit     int
	offset    int
	ids       string
	near      string
	sort      string
	order     string
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

func (repository *rentalRepository) searchRentals(params *rentalSearchParams) ([]*Rental, error) {
	query, args := buildSearchQuery(params)

	rows, err := repository.database.Query(query, args...)
	if err != nil {
		return []*Rental{}, err
	}
	rentals := []*Rental{}

	for rows.Next() {

		rental := Rental{}
		rows.Scan(&rental.Id, &rental.Name, &rental.Description, &rental.Type, &rental.Make, &rental.Model, &rental.Year, &rental.Length, &rental.Sleeps, &rental.PrimaryImageURL,
			&rental.Price.Day, &rental.Location.City, &rental.Location.State, &rental.Location.Zip, &rental.Location.Country, &rental.Location.Latitude, &rental.Location.Longitude,
			&rental.User.Id, &rental.User.FirstName, &rental.User.LastName)
		rentals = append(rentals, &rental)
	}

	return rentals, nil
}

func getIntArrayFromString(intList string) []int {
	stringIDs := strings.Split(intList, ",")

	intIDs := []int{}
	for _, value := range stringIDs {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		intIDs = append(intIDs, intValue)
	}

	return intIDs
}

func getFloatArrayFromString(intList string) []float64 {
	stringIDs := strings.Split(intList, ",")

	floatIDs := []float64{}
	for _, value := range stringIDs {
		floatValue, err := strconv.ParseFloat(value, 32)
		if err != nil {
			panic(err)
		}
		floatIDs = append(floatIDs, floatValue)
	}

	return floatIDs
}

func buildSearchQuery(params *rentalSearchParams) (string, []interface{}) {
	query := "select rentals.id, name, description, type, vehicle_make, vehicle_model, vehicle_year, vehicle_length,  sleeps, primary_image_url, price_per_day, home_city, home_state, home_zip, home_country, lat, lng, u.id, u.first_name, u.last_name from rentals, users u where u.id = user_id"
	args := []interface{}{}

	argCount := 0
	//ids
	if params.ids != "" {
		argCount += 1
		query += fmt.Sprintf(" and rentals.id = ANY ($%d)", argCount)
		args = append(args, pq.Array(getIntArrayFromString(params.ids)))
	}

	//price min
	if params.price_min != 0 {
		argCount += 1
		query += fmt.Sprintf(" and price_per_day >= $%d", argCount)
		args = append(args, params.price_min)
	}

	//price max
	if params.price_max != 0 {
		argCount += 1
		query += fmt.Sprintf(" and price_per_day <= $%d", argCount)
		args = append(args, params.price_max)
	}

	//near
	if params.near != "" {
		point := getFloatArrayFromString(params.near)
		lat := point[0]
		lng := point[1]
		query += fmt.Sprintf(" and ( 3959 * acos( cos( radians(%f) ) * cos( radians( lat ) ) * cos( radians( lng ) - radians(%f) ) + sin( radians(%f) ) * sin( radians( lat ) ) ) ) <= 100", lat, lng, lat)
	}

	//sort
	if params.sort != "" {
		query += fmt.Sprintf(" order by %s %s", params.sort, params.order)
	}

	//pagination
	query += fmt.Sprintf(" OFFSET %d", params.offset)

	if params.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", params.limit)
	}

	return query, args
}
