package service

import (
	"database/sql"
	"fmt"
	"main/bayarXYZ/entity"
	"main/bayarXYZ/request"
	"main/bayarXYZ/response"
	"time"
)

type PaymentService struct {
	DB *sql.DB
}

func NewPaymentService(db *sql.DB) PaymentService {
	return PaymentService{
		DB: db,
	}
}

func (h PaymentService) CreatePayment(req request.CreatePaymentRequest) (int64, error) {

	// init repo
	cRepo := entity.NewCouponRepository(h.DB)
	pRepo := entity.NewPaymentRepository(h.DB)
	vRepo := entity.NewVirtualAccountRepository(h.DB)

	// check source id already mapped
	existPayment, err := pRepo.GetPaymentByPaymentTypeSourceID(req.PaymentType, req.SourceID)
	if existPayment.ID > 0 {
		return 0, fmt.Errorf("CreatePayment: VA %v already mapped", existPayment.ID)
	}
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	}

	var va entity.VirtualAccount
	// check VA
	if req.PaymentType == 1 {
		va, err = vRepo.GetVAByUserIDCode(req.UserID, req.SourceID)
		if err != nil {
			return 0, err
		}
		// VA Expired
		expireVA := time.Unix(va.CreatedAt, 0).Add(10 * time.Hour)
		if time.Now().After(expireVA) {
			return 0, fmt.Errorf("CreatePayment: VA expired")
		}
	}

	var cAmount int32

	// get discount detail
	if req.CouponID > 0 {
		c, err := cRepo.GetCouponByID(req.CouponID)
		if c.ID == 0 {
			return 0, fmt.Errorf("CreatePayment: coupon not found")
		}
		if err != nil {
			return 0, err
		}

		cAmount = c.Amount

		if c.Percentage > 0 {
			cAmount = req.Amount / 100 * c.Percentage
		}
	}

	// add to database
	tNow := time.Now()
	p := entity.Payment{
		UserID:         req.UserID,
		PaymentType:    req.PaymentType,
		Amount:         req.Amount,
		DiscountAmount: cAmount,
		Description:    req.Description,
		SourceID:       req.SourceID,
		CreatedAt:      tNow.Unix(),
		UpdatedAt:      tNow.Unix(),
	}

	id, err := pRepo.Create(p)
	if err != nil {
		return 0, err
	}

	// update va
	if va.ID > 0 {
		err = vRepo.UpdateUsedVA(1, va.ID)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (h PaymentService) ListPaymentVA(pType int8) ([]response.PaymentResponse, error) {

	var result []response.PaymentResponse
	pRepo := entity.NewPaymentRepository(h.DB)

	list, err := pRepo.GetPaymentByPaymentType(pType)
	if err != nil {
		return result, err
	}

	return response.ConstructResponseList(list), err
}
