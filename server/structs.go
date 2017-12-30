package main

import (
	"time"
)

type Todo struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title" binding:"required"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	Completed bool   `db:"completed" json:"completed"`
}

func newTodo(title string) Todo {
	return Todo{
		CreatedAt: time.Now().UnixNano(),
		Title:     title,
		Completed: false,
	}
}

type User struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Address   string `db:"address" json:"address"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

func newUser(name string, address string) User {
	return User{
		CreatedAt: time.Now().UnixNano(),
		Name:      name,
		Address:   address,
	}
}
