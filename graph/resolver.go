//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"sync"

	"github.com/ophum/gqlgen-todos-learn/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	mu        sync.RWMutex
	todos     []*model.Todo
	histories []*model.History
}
