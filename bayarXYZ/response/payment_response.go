package response

import (
	"main/bayarXYZ/entity"
	"time"
)

type PaymentResponse struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	Amount         int32  `json:"amount"`
	DiscountAmount int32  `json:"discount_amount"`
	Description    string `json:"description"`
	SourceID       string `json:"source_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func ConstructResponseList(list []entity.Payment) []PaymentResponse {
	var res []PaymentResponse

	for _, value := range list {
		data := PaymentResponse{
			ID:             value.ID,
			UserID:         value.UserID,
			Amount:         value.Amount,
			DiscountAmount: value.DiscountAmount,
			Description:    value.Description,
			SourceID:       value.SourceID,
			CreatedAt:      time.Unix(value.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:      time.Unix(value.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
		}

		res = append(res, data)
	}

	return res
}
