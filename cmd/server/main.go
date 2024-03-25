package main

import (
	"context"
	"fmt"
	"log"
	"time"

	authModels "github.com/sherpaurgen/Tilicho/internal/auth/models"
	"github.com/sherpaurgen/Tilicho/internal/auth/services"
	"github.com/sherpaurgen/Tilicho/internal/database"
)

// run will instantiate and startup the project app
func Run() error {
	fmt.Println("Starting...")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("cannot connect to database")
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Ping(ctx)
	if err != nil {
		log.Printf("failed to ping database: %v", err)
		return err
	}
	err = db.Migratedb()
	if err != nil {
		err = fmt.Errorf("migratedb failed: %w", err)
		fmt.Println(err)
	}
	fmt.Println("database connection [ok]")

	usersvc := services.NewUserService(db)
	fmt.Println(usersvc.GetUserByUsername(context.Background(), "jane_smith"))
	var v = authModels.User{
		//Userid:   "550e8400-e29b-41d4-a716-446655440000", //his will be auto gen by postgres db
		Username: "john_doe",
		Email:    "john@example.com",
		Password: "password123",
		Groups:   "admin",
		IsActive: "true",
	}
	usersvc.CreateUser(context.Background(), v)
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
