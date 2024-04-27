package request

type CreatePaymentRequest struct {
	UserID      int64  `json:"user_id" binding:"required"`
	SourceID    string `json:"source_id" binding:"required"`
	Amount      int32  `json:"amount" binding:"required"`
	CouponID    int64  `json:"coupon_id"`
	PaymentType int8   `json:"payment_type" binding:"required"`
	Description string `json:"description" binding:"required"`
}
