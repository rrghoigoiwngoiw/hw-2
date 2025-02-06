package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/frrghoigoiwngoiw/hw-2/hw15_go_sql/database"
	"github.com/frrghoigoiwngoiw/hw-2/hw15_go_sql/handlers"
)

func main() {
	db, err := database.NewConnection("dbonlinestore", "db_user", "password", "localhost", 5432)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	h := &handlers.Handlers{DB: db}

	http.HandleFunc("/users/add", h.CreateUserHandler)
	http.HandleFunc("/users/get", h.GetUserHandler)
	http.HandleFunc("/products/add", h.CreateProductHandler)
	http.HandleFunc("/orders", h.GetOrdersHandler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Сервер запущен на :8080")
		if listenErr := server.ListenAndServe(); listenErr != nil && listenErr != http.ErrServerClosed {
			log.Printf("Ошибка запуска сервера: %v", listenErr)
		}
	}()

	<-done
	log.Println("Сервер завершает работу...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if shutdownErr := server.Shutdown(ctx); shutdownErr != nil {
		log.Printf("Ошибка при завершении работы сервера: %v", shutdownErr)
	}

	log.Println("Сервер успешно остановлен")
}
