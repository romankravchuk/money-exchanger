package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next CurrencyConverter
}

func NewLoggingService(next CurrencyConverter) CurrencyConverter {
	return &loggingService{
		next: next,
	}
}

func (s loggingService) Convert(ctx context.Context, from, to string, amount float64) (result float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"time":      time.Since(begin),
			"err":       err,
			"result":    result,
		}).Info("moneyConverter")
	}(time.Now())

	return s.next.Convert(ctx, from, to, amount)
}
