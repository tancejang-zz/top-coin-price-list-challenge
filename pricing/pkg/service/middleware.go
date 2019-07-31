package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(PricingService) PricingService

type loggingMiddleware struct {
	logger log.Logger
	next   PricingService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a PricingService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PricingService) PricingService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetPricingInformation(ctx context.Context, currency string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "GetPricingInformation", "currency", currency, "s0", s0, "e1", e1)
	}()
	return l.next.GetPricingInformation(ctx, currency)
}
