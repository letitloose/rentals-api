package rentals

type RentalService struct {
	repository *rentalRepository
}

func NewRentalService(repository *rentalRepository) *RentalService {
	return &RentalService{repository: repository}
}

func (service *RentalService) GetRental(id int) (*Rental, error) {
	return service.repository.getRental(id), nil
}
