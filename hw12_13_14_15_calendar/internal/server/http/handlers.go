package internalhttp

import (
	"fmt"
	"net/http"

	"github.com/Gilfoyle3301/hw_golang/hw12_13_14_15_calendar/internal/logger"
)

type NewHandler struct {
	logger *logger.Logger
}

func (h *NewHandler) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
