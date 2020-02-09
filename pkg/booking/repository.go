package booking

import (
	"github.com/jmoiron/sqlx"
	"space-trouble/internal/date"
)

type Repository interface {
	Save(booking Booking) error
	GetAll() ([]Booking, error)
	Get(launchpadID string, launchDate date.Date) (b Booking, ok bool)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) Save(b Booking) error {
	_, err :=  r.db.Exec("insert into bookings(first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) " +
		"values (?, ?, ?, ?, ?, ?, ?)", b.FirstName, b.LastName, b.Gender, b.Birthday.Format(date.Format), b.LaunchpadID, b.DestinationID, b.LaunchDate.Format(date.Format))
	return err
}

func (r *repository) GetAll() (all []Booking, err error) {
	rows, err := r.db.Queryx("select * from bookings")
	if err != nil {
		return all, err
	}
	defer rows.Close()
	for rows.Next() {
		var booking Booking
		err := rows.StructScan(&booking)
		if err != nil {
			return all, err
		}
		all = append(all, booking)
	}
	return all, nil
}

func (r *repository) Get(launchpadID string, launchDate date.Date) (b Booking, ok bool) {
	rows, err := r.db.Queryx("select * from bookings where launchpad_id = ? and launch_date = ?", launchpadID, launchDate.Format(date.Format))
	if err != nil {
		return b, false
	}
	defer rows.Close()
	ok = rows.Next()
	err = rows.StructScan(&b)
	if err != nil {
		return b, false
	}
	return b, ok
}
