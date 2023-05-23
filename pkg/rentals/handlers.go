package rentals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (rentalService *RentalService) AddHandlersToMux(mux *http.ServeMux) {
	mux.HandleFunc("/rentals/", rentalService.GetRentalJSON)
	mux.HandleFunc("/rentals", rentalService.SearchRentalsJSON)
}

func extractID(request *http.Request) (int, error) {
	urlPath := strings.Split(request.URL.Path[1:], "/")
	return strconv.Atoi(urlPath[len(urlPath)-1])
}

func extractParams(request *http.Request) (*RentalSearchParams, error) {
	var priceMin, priceMax, limit, offset int
	var err error
	if request.URL.Query().Get("price_min") != "" {
		priceMin, err = strconv.Atoi(request.URL.Query().Get("price_min"))
		if err != nil {
			return nil, err
		}
	}
	if request.URL.Query().Get("price_max") != "" {
		priceMax, err = strconv.Atoi(request.URL.Query().Get("price_max"))
		if err != nil {
			return nil, err
		}
	}
	if request.URL.Query().Get("limit") != "" {
		limit, err = strconv.Atoi(request.URL.Query().Get("limit"))
		if err != nil {
			return nil, err
		}
	}
	if request.URL.Query().Get("offset") != "" {
		limit, err = strconv.Atoi(request.URL.Query().Get("offset"))
		if err != nil {
			return nil, err
		}
	}
	params := &RentalSearchParams{
		Ids:      request.URL.Query().Get("ids"),
		PriceMin: priceMin,
		PriceMax: priceMax,
		Near:     request.URL.Query().Get("near"),
		Sort:     request.URL.Query().Get("sort"),
		Order:    request.URL.Query().Get("order"),
		Limit:    limit,
		Offset:   offset,
	}

	return params, nil
}

func (rentalService *RentalService) GetRentalJSON(writer http.ResponseWriter, request *http.Request) {
	rentalID, err := extractID(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, err.Error())
		return
	}

	rental, err := rentalService.GetRental(rentalID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, err.Error())
		return
	}

	bytes, _ := json.Marshal(rental)
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}

func (rentalService *RentalService) SearchRentalsJSON(writer http.ResponseWriter, request *http.Request) {
	params, err := extractParams(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, err.Error())
		return
	}

	rental, err := rentalService.SearchRentals(params)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, err.Error())
		return
	}

	bytes, _ := json.Marshal(rental)
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}
