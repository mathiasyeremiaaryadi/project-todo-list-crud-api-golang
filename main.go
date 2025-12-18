package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DBConnection *gorm.DB
)

func main() {
	log.Info("Load environment configuration . . .")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment configuration: %v", err)
	}
	log.Info("Success load environment configuration . . .")

	log.Info("Connecting database . . .")
	DBConnection, err = NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	log.Info("Success connecting database")

	log.Info("Migrating all tables . . .")
	MigrateTables(DBConnection)
	log.Info("Success migrating all tables")

	app := fiber.New()

	NewRoute(app)

	go func() {
		applicationHost := fmt.Sprintf("%s:%s", os.Getenv("APPLICATION_HOST"), os.Getenv("APPLICATION_PORT"))
		if err := app.Listen(applicationHost); err != nil {
			log.Fatalf("Failed to start the server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Info("Gracefully shutting down . . .")
	_ = app.Shutdown()

	log.Info("Server was successful shutdown")
}
