package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/frrghoigoiwngoiw/hw-2/hw16_docker/database"
)

type Handlers struct {
	DB *database.DB
}

func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateUser(user.Username, user.Password, user.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Пользователь добавлен")
}

func (h *Handlers) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	user, err := h.DB.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateProduct(product.Name, product.Price); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Продукт добавлен")
}

func (h *Handlers) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	orders, err := h.DB.GetOrdersByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
