package rentals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (rentalService *RentalService) AddHandlersToMux(mux *http.ServeMux) {
	mux.HandleFunc("/users/", rentalService.GetRentalJSON)
}

func extractID(r *http.Request) (int, error) {
	urlPath := strings.Split(r.URL.Path[1:], "/")
	return strconv.Atoi(urlPath[len(urlPath)-1])
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
