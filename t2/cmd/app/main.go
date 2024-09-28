package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"t1/t2/internal/database"
	"t1/t2/internal/handlers"
	"t1/t2/internal/messageService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messageService.Message{})

	repo := messageService.NewMessageRepository(database.DB)
	service := messageService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/message", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/message", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/message", handler.DeleteMessageHandler).Methods("DELETE")
	router.HandleFunc("/api/message", handler.UpdateMessageHandler).Methods("PATCH")

	http.ListenAndServe(":8080", router)
}
