package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ASoldo/golang-ecolabel-backend/internal/config"
	"github.com/ASoldo/golang-ecolabel-backend/internal/handlers"
	"github.com/ASoldo/golang-ecolabel-backend/internal/models"
)

func TestHandleLogin(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedStatus int
	}{
		{
			name:           "Valid credentials",
			input:          `{"username":"` + models.DemoUser.Username + `","password":"` + models.DemoUser.Password + `"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid credentials",
			input:          `{"username":"wrong","password":"wrong"}`,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Insufficient credentials",
			input:          `{"username":"","password":""}`,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid request body",
			input:          `{"username":,}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(test.input))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.HandleLogin)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, test.expectedStatus)
			}
		})
	}
}

func TestHandleDashboard(t *testing.T) {
	token := models.GenerateToken(config.JwtSecret)

	req, err := http.NewRequest("GET", "/dashboard", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleDashboard)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func BenchmarkHandleLogin(b *testing.B) {
	loginReq := `{"username":"test","password":"test"}`
	reqBody := ioutil.NopCloser(bytes.NewReader([]byte(loginReq)))
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest(http.MethodPost, "/login", reqBody)
		handlers.HandleLogin(w, r)
	}
}

func BenchmarkHandleDashboard(b *testing.B) {
	w := httptest.NewRecorder()

	// Generate a valid JWT token to be used in the Authorization header
	token := models.GenerateToken(config.JwtSecret)
	authHeader := "Bearer " + token

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
		r.Header.Set("Authorization", authHeader)
		handlers.HandleDashboard(w, r)
	}
}
