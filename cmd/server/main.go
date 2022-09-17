package main

import (
	"fmt"
	"log"

	"github.com/arekgazdziak/proglog/internal/server"
)

func main() {
	fmt.Println("Starting server...")
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
