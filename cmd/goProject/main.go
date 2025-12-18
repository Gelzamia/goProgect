package main

import (
	"log"
	"net/http"

	"github.com/Gelzamia/goProject/internal/config"
	"github.com/Gelzamia/goProject/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
	cfg := config.Load()
	http.ListenAndServe(":"+cfg.Port, mux)

}
