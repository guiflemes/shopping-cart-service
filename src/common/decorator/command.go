package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func ApplyCommandDecorator[T any](handler CommandHandler[T], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[T] {
	return commandLoggingDecorator[T]{
		base: commandMetricsDecorator[T]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func actionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
