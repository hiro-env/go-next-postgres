package main

import (
	"app/config"
	"app/db"
	"app/handlers"
	"app/repositories"
	"app/routes"
	"app/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	accountRepo := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	e := echo.New() // Echoのインスタンスを作成

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // すべてのオリジンを許可
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	routes.Init(e, userHandler, accountHandler)
	e.Logger.Fatal(e.Start(":8080")) // 8080ポートでサーバを起動
}
