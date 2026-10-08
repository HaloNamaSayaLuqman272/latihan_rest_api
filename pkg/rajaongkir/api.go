package rajaongkir

import (
	"context"
	"encoding/json"
	"latihan_rest_api/pkg/clients"
	"latihan_rest_api/pkg/constant"
	"latihan_rest_api/pkg/utils"
	"net/http"
	"strconv"
)

const BASE_URL = "https://rajaongkir.komerce.id/api/v1"

type Service interface {
	GetDeliveryFee(ctx context.Context, req GetFeeRequest) (float64, string, error)
}

type service struct {
	client clients.HTTPClient
}

func InitService() Service {
	return &service{
		client: clients.InitHTTPClient(BASE_URL, 10, utils.GetConfigurance(constant.RAJAONGKIR_API_KEY)),
	}
}

func (r *service) GetDeliveryFee(ctx context.Context, req GetFeeRequest) (float64, string, error) {
	var response FeeResponse

	payload := map[string]string{
		"origin":      req.Origin,
		"destination": req.Destination,
		"weight":      strconv.Itoa(int(req.Weight)),
		"courier":     req.Courier,
	}

	res, err := r.client.SendFormEncoded(
		ctx,
		"/calculate/district/domestic-cost",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return 0, "", err
	}
	if err := json.Unmarshal([]byte(res), &response); err != nil {
		return 0, "", err
	}

	fee := float64(response.Data[0].Cost)
	etd := response.Data[0].Etd

	return fee, etd, nil
}
