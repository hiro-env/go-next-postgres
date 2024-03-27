package main

import (
	"app/config"
	"app/db"
	"app/handlers"
	"app/repositories"
	"app/routes"
	"app/services"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()
	db, err := db.NewDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New() // Echoのインスタンスを作成
	routes.Init(e, userHandler)
	e.Logger.Fatal(e.Start(":8080")) // 8080ポートでサーバを起動
}
