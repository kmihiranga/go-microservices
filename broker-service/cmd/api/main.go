package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/kmihiranga/broker-service/handlers"
)

const webPort = "80"

func main() {
	app := handler.Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}