package app

import (
	"log"
	"net/http"

	"github.com/MalickSidi/mvc01/controllers"
)

func StartApp() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/users", controllers.GetUser)
	log.Println("Server listenning on port:", server.Addr)
	server.ListenAndServe()
}
