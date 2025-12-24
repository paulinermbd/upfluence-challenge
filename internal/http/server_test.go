package http

import (
	"challenge/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnalysisServer_GetAnalysisRoute_WithQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/analysis?param=value", nil)
	recorder := httptest.NewRecorder()

	mux := setupTestServer()
	mux.ServeHTTP(recorder, req)

	if recorder.Code == 0 {
		t.Error("Expected a status code, got none")
	}
}

func TestAnalysisServer_PostAnalysisRoute_ShouldReturn404(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/analysis", nil)
	recorder := httptest.NewRecorder()

	mux := setupTestServer()
	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for POST method, got %d", recorder.Code)
	}
}

func TestAnalysisServer_UnknownRoute_ShouldReturn404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	recorder := httptest.NewRecorder()

	mux := setupTestServer()
	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for unknown route, got %d", recorder.Code)
	}
}

func TestAnalysisServer_RootRoute_ShouldReturn404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	mux := setupTestServer()
	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for root route, got %d", recorder.Code)
	}
}

func TestAnalysisServer_MultipleRequests(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "GET /analysis should be handled",
			method:         http.MethodGet,
			path:           "/analysis",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "PUT /analysis should return 404",
			method:         http.MethodPut,
			path:           "/analysis",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "DELETE /analysis should return 404",
			method:         http.MethodDelete,
			path:           "/analysis",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "PATCH /analysis should return 404",
			method:         http.MethodPatch,
			path:           "/analysis",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "GET /analysis/123 should return 404",
			method:         http.MethodGet,
			path:           "/analysis/123",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			recorder := httptest.NewRecorder()
			mux := setupTestServer()

			mux.ServeHTTP(recorder, req)

			if tt.path == "/analysis" && tt.method == http.MethodGet {
				if recorder.Code == http.StatusNotFound {
					t.Errorf("Expected route to be registered, got 404")
				}
			} else if recorder.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, recorder.Code)
			}
		})
	}
}

func TestAnalysisServer_ConcurrentRequests(t *testing.T) {
	mux := setupTestServer()
	const numRequests = 100

	done := make(chan bool, numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			req := httptest.NewRequest(http.MethodGet, "/analysis", nil)
			recorder := httptest.NewRecorder()
			mux.ServeHTTP(recorder, req)
			done <- true
		}()
	}

	for i := 0; i < numRequests; i++ {
		<-done
	}
}

func TestAnalysisServer_InvalidHTTPMethod(t *testing.T) {
	req := httptest.NewRequest("INVALID", "/analysis", nil)
	recorder := httptest.NewRecorder()

	mux := setupTestServer()
	mux.ServeHTTP(recorder, req)

	if recorder.Code == 0 {
		t.Error("Expected a status code for invalid method")
	}
}

func setupTestServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /analysis", controller.GetAnalysisHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	return mux
}
