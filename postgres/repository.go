package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GoranP/innsecure"
)

type BookingRepo struct {
	db *sql.DB
}

// NewRepo returns a new repository backed by the given DB.
func NewRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{
		db: db,
	}
}

// Insert satisfies Repository.
func (r *BookingRepo) Insert(ctx context.Context, b innsecure.Booking) error {
	stmt := string("insert into \"Bookings\" (id, hotelid, arrive, leave, name) values ($1, $2, $3, $4, $5)")
	_, err := r.db.Exec(stmt, b.ID, b.HotelID, b.Arrive, b.Leave, b.Name)
	return err
}

// List returns the full contents of the repository.
func (r *BookingRepo) List(ctx context.Context, hotelID int) ([]innsecure.Booking, error) {
	stmt := string("select id, hotelid, arrive, leave, name from \"Bookings\" WHERE hotelid=$1 ")
	rows, err := r.db.Query(stmt, hotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to list bookings: %w", err)
	}
	result := []innsecure.Booking{}
	for rows.Next() {
		var b innsecure.Booking
		err = rows.Scan(&b.ID, &b.HotelID, &b.Arrive, &b.Leave, &b.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result = append(result, b)
	}
	return result, nil
}

// ByID returns a single booking by ID.
// If no booking is found with the given ID, no error is returned.
func (r *BookingRepo) ByID(ctx context.Context, hotelID int, ID string) (*innsecure.Booking, error) {
	//q := fmt.Sprintf("select id, hotelid, arrive, leave, name from Bookings where hotelid='%d' AND id='%s'", hotelID, ID)
	//rows, err := r.db.Query(q)
	q := string("select id, hotelid, arrive, leave, name from \"Bookings\" where hotelid=$1 AND id=$2 ")
	rows, err := r.db.Query(q, hotelID, ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b innsecure.Booking
		err = rows.Scan(&b.ID, &b.HotelID, &b.Arrive, &b.Leave, &b.Name)
		if err == nil {
			return &b, nil
		}
	}
	return nil, nil
}
