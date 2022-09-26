package service

import (
	"context"
	"shopping_cart/src/app"
	"shopping_cart/src/app/command"
	"shopping_cart/src/common/metrics"

	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.FakeMetric{}

	return app.Application{
		Commands: app.Commands{
			AddItem: command.NewAddItemHandler(logger, metricsClient),
		},
	}
}
