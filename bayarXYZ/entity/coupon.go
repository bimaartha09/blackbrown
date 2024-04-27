package entity

import "database/sql"

type Coupon struct {
	ID          int64
	Name        string
	Description string
	Percentage  int32
	Amount      int32
	ExpireTime  int64
	CreatedAt   int64
	UpdatedAt   int64
}

type CouponRepository struct {
	DB *sql.DB
}

func NewCouponRepository(db *sql.DB) CouponRepository {
	return CouponRepository{
		DB: db,
	}
}

func (r CouponRepository) GetCouponByID(id int64) (Coupon, error) {
	var c Coupon

	row := r.DB.QueryRow("SELECT * FROM coupon WHERE id = ?", id)
	if err := row.Scan(&c.ID, &c.Name, &c.Description, &c.Percentage,
		&c.Amount, &c.ExpireTime, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return c, err
	}

	return c, nil
}
