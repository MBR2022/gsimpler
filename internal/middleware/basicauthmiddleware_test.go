package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MBR2022/gosimpler/internal/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	username = "admin"
	password = "admin"
)

func TestBasicAuthFailed(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.SetBasicAuth(username, "wrong-password")
	handler := middleware.NewBasicAuthMiddleware(username, password).Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "test")
		_, err := w.Write([]byte("content"))
		assert.Nil(t, err)

		flusher, ok := w.(http.Flusher)
		assert.True(t, ok)
		flusher.Flush()

	})
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestBasicAuthSuccess(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.SetBasicAuth(username, password)
	handler := middleware.NewBasicAuthMiddleware(username, password).Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "test")
		_, err := w.Write([]byte("content"))
		assert.Nil(t, err)

		flusher, ok := w.(http.Flusher)
		assert.True(t, ok)
		flusher.Flush()

	})
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "test", resp.Header().Get("X-Test"))
	assert.Equal(t, "content", resp.Body.String())
}
