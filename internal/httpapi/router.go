package httpapi

import (
    "encoding/json"
    "net/http"
    "time"
)

type statusResponse struct {
    Service string `json:"service"`
    Status  string `json:"status"`
    Version string `json:"version"`
    Time    string `json:"time"`
}

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, http.StatusOK, statusResponse{
            Service: "cherryshare-api",
            Status:  "ok",
            Version: "0.1.0-dev",
            Time:    time.Now().UTC().Format(time.RFC3339),
        })
    })

    mux.HandleFunc("GET /api/v1", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, http.StatusOK, map[string]any{
            "name":    "CherryShare API",
            "version": "v1",
            "features": []string{
                "files",
                "shares",
                "versions",
                "tenants",
                "audit",
                "ai-search-planned",
            },
        })
    })

    return requestID(securityHeaders(mux))
}

func writeJSON(w http.ResponseWriter, status int, value any) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(value)
}

func securityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("Referrer-Policy", "no-referrer")
        w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
        next.ServeHTTP(w, r)
    })
}

func requestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := r.Header.Get("X-Request-ID")
        if id == "" {
            id = time.Now().UTC().Format("20060102T150405.000000000")
        }
        w.Header().Set("X-Request-ID", id)
        next.ServeHTTP(w, r)
    })
}
