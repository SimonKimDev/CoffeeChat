package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		status int
		body   string
	}{
		{"health", http.MethodGet, "/healthz", http.StatusOK, "{\"status\":\"ok\"}\n"},
		{"head", http.MethodHead, "/healthz", http.StatusOK, ""},
		{"unknown route", http.MethodGet, "/messages", http.StatusNotFound, "404 page not found\n"},
		{"unsupported method", http.MethodPost, "/healthz", http.StatusMethodNotAllowed, "Method Not Allowed\n"},
	}
	srv := httptest.NewServer(NewHandler())
	defer srv.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, srv.URL+tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := srv.Client().Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Errorf("status = %d, want %d", resp.StatusCode, tt.status)
			}
			if tt.status == http.StatusOK && resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Content-Type = %q, want application/json", resp.Header.Get("Content-Type"))
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != tt.body {
				t.Errorf("body = %q, want %q", body, tt.body)
			}
		})
	}
}
