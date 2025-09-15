package main

import (
	"eBlog/internal/configs"
	"eBlog/internal/controller"
	"eBlog/internal/db"
	"eBlog/internal/repository"
	"eBlog/internal/service"
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

	dbConn, err := db.InitConnection()
	if err != nil {
		fmt.Println("Error during database connection initialization: ", err.Error())
		return
	}

	if err = db.RunMigrations(dbConn); err != nil {
		fmt.Println("Error during database migrations: ", err.Error())
		return
	}

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	if err = ctrl.InitRoutes(); err != nil {
		fmt.Println("Error during http-service initialization: ", err.Error())
		return
	}

	if err = db.CloseConnection(dbConn); err != nil {
		fmt.Println("Error during database connection close: ", err.Error())
		return
	}

	fmt.Println("Starting up - End")
}
