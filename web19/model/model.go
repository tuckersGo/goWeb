package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(name string) *Todo
	removeTodo(id int) bool
	completeTodo(id int, complete bool) bool
}

var handler dbHandler

func init() {
	//handler = newMemoryHandler()
	handler = newSqliteHandler()
}

func GetTodos() []*Todo {
	return handler.getTodos()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.completeTodo(id, complete)
}
