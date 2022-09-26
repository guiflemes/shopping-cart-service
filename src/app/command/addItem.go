package command

import (
	"context"
	"errors"
	"shopping_cart/src/common/decorator"

	"github.com/sirupsen/logrus"
)

type AddItemCommand struct {
	CartId   string
	ItemId   string
	Quantity int32
}

type AddItemHandler decorator.CommandHandler[AddItemCommand]

type addItemHandler struct{}

func NewAddItemHandler(logger *logrus.Entry, metricsClient decorator.MetricsClient) AddItemHandler {
	return decorator.ApplyCommandDecorator[AddItemCommand](
		addItemHandler{},
		logger,
		metricsClient,
	)
}

func (c addItemHandler) Handle(ctx context.Context, cmd AddItemCommand) error {
	return errors.New("any error")
}
