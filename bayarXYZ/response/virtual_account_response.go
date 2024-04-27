package response

import (
	"main/bayarXYZ/entity"
)

type VirtualAccountResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

func ConstructVAResponse(obj entity.VirtualAccount) VirtualAccountResponse {
	return VirtualAccountResponse{
		ID:   obj.ID,
		Code: obj.Code,
	}
}
