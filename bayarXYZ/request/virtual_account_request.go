package request

type GenerateVARequest struct {
	UserID string `json:"user_id" binding:"required"`
}
