package main

import (
	"fmt"
	"test-books-wishlist/cmd/config"
	"test-books-wishlist/internal/handler"
	"test-books-wishlist/internal/repository"
	"test-books-wishlist/internal/service"
)

// @title Google Books API - TESTCASE
// @version 1.0
// @description Google Books testcase API
// @BasePath /
func main() {
	app := service.NewService()
	app.Repo = repository.NewPostgresRepository()

	go initDB(app.Repo)

	r := handler.InitRouter(app)

	err := r.Run(fmt.Sprintf(":%s", config.Config.Port))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func initDB(repo repository.Repository) error {
	c := config.Config
	if err := repo.Init(c.DBIP, c.DBPort, c.DBUser, c.DBPass, c.DBName, c.DBRetryCount); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
