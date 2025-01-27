package server // с сервером вроде все

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Server() {
	address := flag.String("address", "127.0.0.1", "Адрес сервера")
	port := flag.String("port", "8080", "Порт сервера")

	flag.Parse()

	serverAddr := fmt.Sprintf("%s:%s", *address, *port)

	fmt.Printf("Сервер запущен на %s\n", serverAddr)

	//  маршруты сервера
	http.HandleFunc("/get", HandleGet)
	http.HandleFunc("/post", HandlePost)

	// Создаем сервер с тайм-аутами
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запускаем сервер
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не подходит", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Зарегистрирован GET-запрос")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "вот ответ и вот еще ответ")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не подходит", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Получен POST-запрос")

	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Данные от клиента: %s\n", string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Клиент говорит: %s", string(body))))
}
