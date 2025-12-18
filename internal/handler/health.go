package handler

import (
	"net/http"

	"github.com/Gelzamia/goProject/internal/service"
)

func Health(w http.ResponseWriter, r *http.Request) {
	result := service.Check()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
