//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"errors"
	"slices"
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

func (r *Resolver) copyTodo(todo *model.Todo) *model.Todo {
	return &model.Todo{
		ID:     todo.ID,
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: todo.UserID,
	}
}

func (r *Resolver) copyHistory(history *model.History) *model.History {
	return &model.History{
		TodoID:    history.TodoID,
		Text:      history.Text,
		Done:      history.Done,
		CreatedAt: history.CreatedAt,
	}

}

func (r *Resolver) addTodo(todo *model.Todo) {
	r.todos = append(r.todos, r.copyTodo(todo))
}

func (r *Resolver) addHistory(history *model.History) {
	r.histories = append(r.histories, r.copyHistory(history))
}

func (r *Resolver) updateTodo(todo *model.Todo) error {
	index := r.todoIndex(todo.ID)
	if index == -1 {
		return errors.New("not found")
	}
	r.todos[index] = r.copyTodo(todo)
	return nil
}

func (r *Resolver) findTodo(id string) (*model.Todo, error) {
	index := r.todoIndex(id)
	if index == -1 {
		return nil, errors.New("not found")
	}
	return r.copyTodo(r.todos[index]), nil
}

func (r *Resolver) todoIndex(id string) int {
	return slices.IndexFunc(r.todos, func(t *model.Todo) bool {
		return t.ID == id
	})
}
