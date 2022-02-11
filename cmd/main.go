package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dreddsa5dies/httprestapient/config"

	"github.com/dreddsa5dies/httprestapient/router"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	//set the client to the variable defined in package config
	//this will enable the client instance to be accessed anywhere through the accessor which is a function
	//named GetClient
	config.SetClient(client)

	//initiate router and register all the route
	r := mux.NewRouter()
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}
