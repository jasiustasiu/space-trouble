package booking

type Repository struct {

}

func (r *Repository) Save(booking Booking) error {
	return nil
}

func (r *Repository) GetAll() ([]Booking, error) {
	return nil, nil
}
