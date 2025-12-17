package main

import "time"

type User struct {
	ID        int `gorm:"primaryKey,autoIncrement"`
	Email     string
	Password  string
	Todo      []Todo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todo struct {
	ID          int `gorm:"primaryKey,autoIncrement"`
	Title       string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
