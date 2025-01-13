package api

import (
	"time"

	"github.com/Paprec/trucktrack/service"
	"github.com/go-kit/kit/metrics"
)

const (
	methode = "method"
)

type metricsMiddleware struct {
	counter metrics.Counter
	latency metrics.Histogram
	service service.MACService
}

func MetricsMiddleware(service service.MACService, counter metrics.Counter, latency metrics.Histogram) service.MACService {
	return &metricsMiddleware{
		counter: counter,
		latency: latency,
		service: service,
	}
}

func (mm *metricsMiddleware) GetMACAddresses(macs []string) []string {
	defer func(begin time.Time) {
		mm.counter.With(methode, EndPointURL).Add(1)
		mm.latency.With(methode, EndPointURL).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.service.GetMACAddresses(macs)
}

func (mm *metricsMiddleware) AuthorId(addmacs string) string {
	defer func(begin time.Time) {
		mm.counter.With(methode, EndPointURL).Add(1)
		mm.latency.With(methode, EndPointURL).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.service.AuthorId(addmacs)
}
