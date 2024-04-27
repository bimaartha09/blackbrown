package service

import (
	"database/sql"
	"main/bayarXYZ/entity"
	"main/bayarXYZ/request"
	"main/bayarXYZ/response"
	"math/rand"
	"strconv"
	"time"
)

type VirtualAccountService struct {
	DB *sql.DB
}

func NewVirtualAccountService(db *sql.DB) VirtualAccountService {
	return VirtualAccountService{
		DB: db,
	}
}

func (h VirtualAccountService) CreateVA(req request.GenerateVARequest) (response.VirtualAccountResponse, error) {

	repo := entity.NewVirtualAccountRepository(h.DB)
	tnow := time.Now()
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return response.VirtualAccountResponse{}, err
	}

	existVA, errLatest := repo.GetLatestAvailableVAByUserID(uid, tnow.Add(time.Duration(-10)*time.Hour).Unix())
	if existVA.ID > 0 {
		return response.ConstructVAResponse(existVA), nil
	}
	if errLatest != nil {
		// pass if error no rows
		if errLatest != sql.ErrNoRows {
			return response.VirtualAccountResponse{}, errLatest
		}
	}

	// get field value
	tnowStr := strconv.Itoa(int(tnow.Unix()))
	randomStr := getDigitRandom(5)
	value := tnowStr + "" + randomStr

	if err != nil {
		return response.VirtualAccountResponse{}, err
	}

	// construct field to input database
	va := entity.VirtualAccount{
		Code:      value,
		UserID:    uid,
		CreatedAt: tnow.Unix(),
	}

	id, err := repo.Create(va)
	if err != nil {
		return response.VirtualAccountResponse{}, err
	}
	va.ID = id

	return response.ConstructVAResponse(va), nil
}

func getDigitRandom(n int) string {
	var str string

	for i := 0; i < n; i++ {
		num := rand.Intn(10)
		str = str + strconv.Itoa(num)
	}

	return str
}
