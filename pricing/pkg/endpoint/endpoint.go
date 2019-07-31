package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/tancejang/go_challenge_pricing_service/pricing/pkg/service"
)

// GetPricingInformationRequest collects the request parameters for the GetPricingInformation method.
type GetPricingInformationRequest struct {
	Currency string `json:"currency"`
}

// GetPricingInformationResponse collects the response parameters for the GetPricingInformation method.
type GetPricingInformationResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeGetPricingInformationEndpoint returns an endpoint that invokes GetPricingInformation on the service.
func MakeGetPricingInformationEndpoint(s service.PricingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPricingInformationRequest)
		s0, e1 := s.GetPricingInformation(ctx, req.Currency)
		return GetPricingInformationResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetPricingInformationResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetPricingInformation implements Service. Primarily useful in a client.
func (e Endpoints) GetPricingInformation(ctx context.Context, currency string) (s0 string, e1 error) {
	request := GetPricingInformationRequest{Currency: currency}
	response, err := e.GetPricingInformationEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetPricingInformationResponse).S0, response.(GetPricingInformationResponse).E1
}
