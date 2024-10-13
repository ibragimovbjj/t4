package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"t1/internal/database"
	"t1/internal/handlers"
	"t1/internal/messageService"
	"t1/internal/web/message"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messageService.Message{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	repo := messageService.NewMessageRepository(database.DB)
	service := messageService.NewService(repo)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := message.NewStrictHandler(handler, nil) // тут будет ошибка
	message.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
