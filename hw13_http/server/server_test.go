package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/get", nil)
	w := httptest.NewRecorder()

	HandleGet(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("ожидался статус %d, получен %d", http.StatusOK, res.StatusCode)
	}

	// Проверяем ответ
	body, _ := io.ReadAll(res.Body)
	expectedBody := "вот ответ и вот еще ответ"
	if string(body) != expectedBody {
		t.Errorf("ожидался ответ '%s', получен '%s'", expectedBody, string(body))
	}
}

func TestHandleGet_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/get", nil)
	w := httptest.NewRecorder()

	HandleGet(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("ожидался статус %d, получен %d", http.StatusMethodNotAllowed, res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	expectedBody := "Метод не подходит\n"
	if string(body) != expectedBody {
		t.Errorf("ожидался ответ '%s', получен '%s'", expectedBody, string(body))
	}
}

func TestHandlePost(t *testing.T) {
	payload := []byte("привет сервер")
	req := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(payload))
	req.ContentLength = int64(len(payload))
	w := httptest.NewRecorder()

	HandlePost(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("ожидался статус %d, получен %d", http.StatusOK, res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	expectedBody := "Клиент говорит: привет сервер"
	if string(body) != expectedBody {
		t.Errorf("ожидался ответ '%s', получен '%s'", expectedBody, string(body))
	}
}

func TestHandlePost_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/post", nil)
	w := httptest.NewRecorder()

	HandlePost(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("ожидался статус %d, получен %d", http.StatusMethodNotAllowed, res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	expectedBody := "Метод не подходит\n"
	if string(body) != expectedBody {
		t.Errorf("ожидался ответ '%s', получен '%s'", expectedBody, string(body))
	}
}
