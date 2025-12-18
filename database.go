package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SCHEMA"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return db, err
	}

	return db, err
}

func MigrateTables(databaseConnection *gorm.DB) {
	databaseConnection.AutoMigrate(
		&User{},
		&Todo{},
	)
}

func CreateUser(user User) error {
	err := DBConnection.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUser(email string) (User, error) {
	var user User
	err := DBConnection.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateTodo(todo Todo) (Todo, error) {
	err := DBConnection.Create(&todo).Error
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func GetTodo(todoId int) (Todo, error) {
	var todo Todo
	err := DBConnection.Where("ID = ?", todoId).First(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func UpdateTodo(todo Todo) (Todo, error) {
	err := DBConnection.Save(&todo).Error
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}
