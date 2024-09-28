package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/themellowpea/godo/internal/db"
	"github.com/themellowpea/godo/internal/server"
)

func init() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "the server address")
	flag.Parse()

	dbPool, err := db.NewGoDoDB().GetDBConnection()
	if err != nil {
		log.Fatalf("Failed to initilize database: %v", err)
	}

	defer dbPool.Close()

	r := gin.Default()
	server := server.NewServer(*listenAddr, r, dbPool)
	if err := server.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
