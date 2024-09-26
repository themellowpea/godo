package main

import (
	"flag"
	"fmt"

	"github.com/themellowpea/godo/internal/server"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "the server address")
	flag.Parse()

	router := server.NewRouter()
	server := server.NewServer(*listenAddr, router)
	err := server.Run()
	if err != nil {
		fmt.Println("Yikes!")
	}
}
