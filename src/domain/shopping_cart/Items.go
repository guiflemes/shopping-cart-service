package shopping_cart

import "github.com/pkg/errors"

type item struct {
	Id       string
	Quantity int32
}

func newItem(itemId string, quantity int32) (item, error) {
	if quantity == 0 {
		return item{}, errors.New("Quantity must be greater than zero")
	}

	return item{itemId, quantity}, nil
}

type Items []item

func (items Items) HasItem(itemId string) bool {
	for _, i := range items {
		if i.Id == itemId {
			return true
		}
	}
	return false
}

func (items *Items) AddItem(item item) {
	*items = append(*items, item)
}

func (items Items) IsEmpty() bool {
	return len(items) == 0
}

func (items Items) ItemCount(itemId string) int {
	count := 0

	for _, item := range items {
		if item.Id == itemId {
			count += 1
		}
	}

	return count
}
