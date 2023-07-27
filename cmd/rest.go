package cmd

import (
	"fmt"
	"go-rest/config"
	"go-rest/database"
	"go-rest/repo"
	"go-rest/rest"
	"go-rest/svc"
)

func serveRest() {
	appConfig := config.GetApp()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", appConfig.DBHost, appConfig.DBUser, appConfig.DBPass, appConfig.DBName, appConfig.DBPort)
	db := database.NewDatabase(dsn)
	userRepo := repo.NewUserRepo(db)
	svc := svc.NewService(userRepo)

	server, err := rest.NewServer(svc, appConfig)

	if err != nil {
		panic(err)
	}

	err = server.Start()

	if err != nil {
		panic(err)
	}
}
