package command

import (
	"context"
	"errors"
	"shopping_cart/src/common/decorator"

	"github.com/sirupsen/logrus"
)

type AddItem struct {
	CartId   string
	ItemId   string
	Quantity int32
}

type AddItemHandler decorator.CommandHandler[AddItem]

type addItemHandler struct{}

func NewAddItemHandler(logger *logrus.Entry, metricsClient decorator.MetricsClient) AddItemHandler {
	return decorator.ApplyCommandDecorator[AddItem](
		addItemHandler{},
		logger,
		metricsClient,
	)
}

func (c addItemHandler) Handle(ctx context.Context, cmd AddItem) error {
	return errors.New("any error")
}
