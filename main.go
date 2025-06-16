package main

import (
	"log"

	"github.com/morigami8/proglog/server"
)

func main() {
	srv := server.NewHTTPServer(":8081")
	log.Fatal(srv.ListenAndServe())
}
