package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const webPort = "80"

func main() {
	app := Config{}
	log.Println("the broker server started at : ", webPort)

	//defined teh server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
