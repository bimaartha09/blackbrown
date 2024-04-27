package entity

import (
	"database/sql"
	"fmt"
)

type VirtualAccount struct {
	ID        int64
	UserID    int64
	Code      string
	IsUsed    int8
	CreatedAt int64
}

type VirtualAccountRepository struct {
	DB *sql.DB
}

func NewVirtualAccountRepository(db *sql.DB) VirtualAccountRepository {
	return VirtualAccountRepository{
		DB: db,
	}
}

func (r VirtualAccountRepository) Create(va VirtualAccount) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO virtual_account (user_id, code, created_at) VALUES (?, ?, ?)", va.UserID, va.Code, va.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("addVirtualAccount: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addVirtualAccount: %v", err)
	}
	return id, nil
}

func (r VirtualAccountRepository) GetLatestAvailableVAByUserID(uid int64, time int64) (VirtualAccount, error) {
	var va VirtualAccount

	row := r.DB.QueryRow("SELECT * FROM virtual_account WHERE user_id = ? AND created_at > ? AND is_used = 0 ORDER BY ID DESC", uid, time)
	if err := row.Scan(&va.ID, &va.UserID, &va.Code, &va.IsUsed, &va.CreatedAt); err != nil {
		return va, err
	}
	return va, nil
}

func (r VirtualAccountRepository) GetVAByCode(code string) (VirtualAccount, error) {
	var va VirtualAccount

	row := r.DB.QueryRow("SELECT * FROM virtual_account WHERE code = ? ORDER BY ID DESC", code)
	if err := row.Scan(&va.ID, &va.Code, &va.Code, &va.IsUsed, &va.CreatedAt); err != nil {
		return va, err
	}

	return va, nil
}

func (r VirtualAccountRepository) GetVAByUserIDCode(uid int64, code string) (VirtualAccount, error) {
	var va VirtualAccount

	row := r.DB.QueryRow("SELECT * FROM virtual_account "+
		"WHERE user_id = ? AND code = ? ORDER BY ID DESC", uid, code)

	if err := row.Scan(&va.ID, &va.Code, &va.Code, &va.IsUsed, &va.CreatedAt); err != nil {
		return va, err
	}

	return va, nil
}

func (r VirtualAccountRepository) UpdateUsedVA(is_used int8, id int64) error {
	_, err := r.DB.Exec("UPDATE virtual_account SET is_used = ? WHERE id = ?", is_used, id)
	if err != nil {
		return err
	}

	return nil
}
