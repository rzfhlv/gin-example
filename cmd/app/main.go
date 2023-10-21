package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal"
	"github.com/rzfhlv/gin-example/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Init()

	svc := internal.New(cfg)

	router := routes.ListRoutes(svc)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cfg.MySQL.Close()
		cfg.Redis.Close()
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 Seconds.")
	}
	log.Println("Server Exiting")
}
