package main

import (
	"context"
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Service Shutdown: ", err)
	}
	log.Println("Service is Shutdown ")
}
