package api

import (
	"fmt"
	"time"

	"github.com/Paprec/trucktrack/service"
	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger  log.Logger
	service service.MACService
}

func LoggingMiddleware(svc service.MACService, logger log.Logger) service.MACService {
	return &loggingMiddleware{
		logger:  logger,
		service: svc,
	}
}

func (lm loggingMiddleware) GetMACAddresses(macs []string) []string {
	defer func(begin time.Time) {
		message := fmt.Sprintf("executing the method took %s to complete", time.Since(begin))
		if len(macs) == 0 {
			lm.logger.Log("message", message, "with error")
		}
		lm.logger.Log("message without errors", message)
	}(time.Now())
	return lm.service.GetMACAddresses(macs)
}

func (lm loggingMiddleware) AuthorId(addmac string) string {
	defer func(begin time.Time) {
		message := fmt.Sprintf("executing the method took %s to complete", time.Since(begin))
		if len(addmac) == 0 {
			lm.logger.Log("message", message, "with error")
		}
		lm.logger.Log("message without errors", message)
	}(time.Now())
	return lm.service.AuthorId(addmac)
}

func (lm loggingMiddleware) PostActivity(addmac string) string {
	defer func(begin time.Time) {
		message := fmt.Sprintf("executing the method took %s to complete", time.Since(begin))
		if len(addmac) == 0 {
			lm.logger.Log("message", message, "with error")
		}
		lm.logger.Log("message without errors", message)
	}(time.Now())
	return lm.service.AuthorId(addmac)
}
