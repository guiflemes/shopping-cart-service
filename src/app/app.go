package app

import "shopping_cart/src/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddItem command.AddItemHandler
}

type Queries struct{}
