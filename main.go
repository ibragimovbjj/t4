package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Msg struct {
	Message string
}

var message string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(message)
	fmt.Fprintln(w, "Hello, World!", message)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var msg Msg
	err := json.NewDecoder(r.Body).Decode(&msg)
	message = msg.Message
	fmt.Println(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintln(w, message)
	}
}

func main() {
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
