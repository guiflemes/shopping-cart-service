package main

import (
	"context"
	"shopping_cart/src/service"
)

func main() {

	ctx := context.Background()

	service.NewApplication(ctx)

}
