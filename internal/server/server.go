package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	listenAddr string
	engine     *gin.Engine
	db         *pgxpool.Pool
}

func New(listenAddr string, engine *gin.Engine, db *pgxpool.Pool) *Server {
	return &Server{
		listenAddr: listenAddr,
		engine:     engine,
		db:         db,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    s.listenAddr,
		Handler: s.engine,
	}

	// Register all the routes
	SetupRoutes(s.engine, s.db)

	// Start server in a goreoutine to allow for graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen error: %v", err)
		}
	}()

	// Handle OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("\nShutting down server...")

	// Create a deadline for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close the db connection pool
	s.db.Close()

	log.Println("\nServer exited gracefully.")
	return nil
}
