package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository interface {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}