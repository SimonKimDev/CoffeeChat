package http

import (
	"encoding/json"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"net/http"
)

type GreetingHandler struct {
	greeter application.Greeter
}

func NewGreetingHandler(g application.Greeter) *GreetingHandler {
	return &GreetingHandler{greeter: g}
}

func (h *GreetingHandler) greet(w http.ResponseWriter, r *http.Request) {
	greeting := h.greeter.Greet()
	_ = json.NewEncoder(w).Encode(greeting)
}
