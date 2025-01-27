package client

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func RunClient() error {
	method := flag.String("method", "GET", "HTTP метод: GET или POST")
	url := flag.String("url", "", "Полный URL сервера")
	data := flag.String("data", "", "Данные для отправки в POST-запросе")
	flag.Parse()

	if *url == "" {
		fmt.Println("Ошибка: Параметр 'url' обязателен")
		flag.Usage()
		os.Exit(1)
	}

	// контекст тайм-аут
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := &http.Client{}

	// Выполнение запроса
	switch *method {
	case "GET":
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, *url, nil)
		if err != nil {
			fmt.Printf("Ошибка при создании GET-запроса: %v\n", err)
			return err
		}

		response, err := client.Do(req)
		if err != nil {
			fmt.Printf("Ошибка при выполнении GET-запроса: %v\n", err)
			return err
		}
		defer response.Body.Close()

		body, _ := io.ReadAll(response.Body)
		fmt.Printf("Ответ сервера: %s\n", string(body))

	case "POST":
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, *url, bytes.NewBufferString(*data))
		if err != nil {
			fmt.Printf("Ошибка при создании POST-запроса: %v\n", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		response, err := client.Do(req)
		if err != nil {
			fmt.Printf("Ошибка при выполнении POST-запроса: %v\n", err)
			return err
		}
		defer response.Body.Close()

		body, _ := io.ReadAll(response.Body)
		fmt.Printf("Ответ сервера: %s\n", string(body))

	default:
		return fmt.Errorf("ошибка: Поддерживаются только методы GET и POST")
	}

	return nil
}
