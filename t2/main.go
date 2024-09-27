package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	err := DB.AutoMigrate(&Message{})
	if err != nil {
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	DB.Create(&msg)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {

	var messages []Message
	DB.Find(&messages)
	err := json.NewEncoder(w).Encode(messages)
	if err != nil {
		return
	}

}
