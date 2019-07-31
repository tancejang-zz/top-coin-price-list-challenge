package service

import "context"

// PricingService describes the service.
type PricingService interface {
	GetPricingInformation(ctx context.Context, currency string) (string, error)
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
}

type basicPricingService struct{}

func (b *basicPricingService) GetPricingInformation(ctx context.Context, currency string) (s0 string, e1 error) {
	// TODO implement the business logic of GetPricingInformation
	return s0, e1
}

// NewBasicPricingService returns a naive, stateless implementation of PricingService.
func NewBasicPricingService() PricingService {
	return &basicPricingService{}
}

// New returns a PricingService with all of the expected middleware wired in.
func New(middleware []Middleware) PricingService {
	var svc PricingService = NewBasicPricingService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
