package booking

import "space-trouble/internal/date"

type Repository struct {

}

func (r *Repository) Save(booking Booking) error {
	return nil
}

func (r *Repository) GetAll() ([]Booking, error) {
	return nil, nil
}

func (r *Repository) Get(launchpadID string, launchDate date.Date) (*Booking, error) {
	return nil, nil
}
