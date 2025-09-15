package main

import (
	"eBlog/internal/configs"
	"eBlog/internal/controller"
	"eBlog/internal/db"
	"fmt"
)

// @title eBlog service
// @contact.name API eBlog
// @contact.url https://test.com/
// @contact.email nekruzrakhimov@icloud.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("Starting up - Start")
	if err := configs.ReadSettings(); err != nil {
		fmt.Println("Error reading settings", err.Error())
		return
	}

	if err := db.InitConnection(); err != nil {
		fmt.Println("Error during database connection initialization: ", err.Error())
		return
	}

	if err := db.RunMigrations(); err != nil {
		fmt.Println("Error during database migrations: ", err.Error())
		return
	}

	if err := controller.InitRoutes(); err != nil {
		fmt.Println("Error during http-service initialization: ", err.Error())
		return
	}

	if err := db.CloseConnection(); err != nil {
		fmt.Println("Error during database connection close: ", err.Error())
		return
	}

	fmt.Println("Starting up - End")
}
