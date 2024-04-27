package entity

import (
	"database/sql"
	"fmt"
)

type Payment struct {
	ID             int64
	UserID         int64
	PaymentType    int8
	Amount         int32
	DiscountAmount int32
	Description    string
	SourceID       string
	CreatedAt      int64
	UpdatedAt      int64
}

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return PaymentRepository{
		DB: db,
	}
}

func (r PaymentRepository) Create(p Payment) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO payment "+
		"(user_id, payment_type, amount, discount_amount, description, source_id, created_at, updated_at) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)", p.UserID, p.PaymentType, p.Amount, p.DiscountAmount, p.Description,
		p.SourceID, p.CreatedAt, p.UpdatedAt)

	if err != nil {
		return 0, fmt.Errorf("addPayment: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addPayment: %v", err)
	}
	return id, nil
}

func (r PaymentRepository) GetPaymentByPaymentTypeSourceID(pType int8, sid string) (Payment, error) {
	var p Payment

	row := r.DB.QueryRow("SELECT * FROM payment WHERE "+
		" payment_type = ? AND source_id = ? ", pType, sid)

	if err := row.Scan(&p.ID, &p.UserID, &p.PaymentType, &p.Amount,
		&p.DiscountAmount, &p.Description, &p.SourceID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return p, err
	}

	return p, nil
}

func (r PaymentRepository) GetPaymentByPaymentType(pType int8) ([]Payment, error) {
	var pList []Payment

	rows, err := r.DB.Query("SELECT * FROM payment WHERE payment_type = ?", pType)

	if err != nil {
		return nil, fmt.Errorf("GetPaymentByPaymentType %q: %v", pType, err)
	}

	defer rows.Close()
	for rows.Next() {
		var p Payment
		if err := rows.Scan(&p.ID, &p.UserID, &p.PaymentType, &p.Amount,
			&p.DiscountAmount, &p.SourceID, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("GetPaymentByPaymentType %q: %v", pType, err)
		}
		pList = append(pList, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPaymentByPaymentType %q: %v", pType, err)
	}

	return pList, nil
}
