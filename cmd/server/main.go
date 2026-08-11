package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/paddman/CherryShare/internal/httpapi"
)

func main() {
    addr := env("CHERRYSHARE_ADDR", ":8080")

    srv := &http.Server{
        Addr:              addr,
        Handler:           httpapi.NewRouter(),
        ReadHeaderTimeout: 5 * time.Second,
        ReadTimeout:       30 * time.Second,
        WriteTimeout:      30 * time.Second,
        IdleTimeout:       60 * time.Second,
    }

    log.Printf("CherryShare API listening on %s", addr)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}

func env(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
