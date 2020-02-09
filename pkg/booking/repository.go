package booking

import (
	"github.com/jmoiron/sqlx"
	"space-trouble/internal/date"
)

const noRowsError = "sql: no rows in result set"

type Repository interface {
	Save(booking Booking) error
	GetAll() ([]Booking, error)
	IsLaunchpadAvailable(launchpadID string, launchDate date.Date, out chan<- AvailabilityResponse)
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
	_, err := r.db.Exec("insert into bookings(first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) "+
		"values ($1, $2, $3, $4, $5, $6, $7)", b.FirstName, b.LastName, b.Gender, b.Birthday.Format(date.Format), b.LaunchpadID, b.DestinationID, b.LaunchDate.Format(date.Format))
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

func (r *repository) IsLaunchpadAvailable(launchpadID string, launchDate date.Date, out chan<- AvailabilityResponse) {
	row := r.db.QueryRowx("select id from bookings where launchpad_id = $1 and launch_date = $2", launchpadID, launchDate.Format(date.Format))
	var id int64
	err := row.Scan(&id)
	if err != nil {
		if err.Error() == noRowsError {
			out <- AvailabilityResponse{Available: true, Err: nil,}
		} else {
			out <- AvailabilityResponse{Available: false, Err: err,}
		}
	} else {
		out <- AvailabilityResponse{Available: false, Err: nil,}
	}
}
