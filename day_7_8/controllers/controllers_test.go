package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"golang-training/day_7_8/controllers"
	"golang-training/day_7_8/models"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	store  *models.LibraryStore
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	store = models.NewBookStore()

	router = gin.New()
	router.POST("/add", func(c *gin.Context) { controllers.AddBookController(c, store) })
	router.DELETE("/delete", func(c *gin.Context) { controllers.DeleteBookController(c, store) })
	router.GET("/books", func(c *gin.Context) { controllers.GetBooksController(c, store) })
	router.PUT("/update", func(c *gin.Context) { controllers.UpdateBookController(c, store) })

	os.Exit(m.Run())
}

func TestAddBookController(t *testing.T) {
	testCases := []struct {
		name       string
		body       string
		statusCode int
	}{
		{
			name:       "Valid Book",
			body:       `{"title":"New Book","author":"John","genre":"Tech","rating":4.5}`,
			statusCode: http.StatusCreated,
		},
		{
			name:       "Missing Title",
			body:       `{"author":"John","genre":"Tech","rating":4.5}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Invalid body",
			body:       `{"title":123,"author":"John","genre":"Tech","rating":4.5}`,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/add", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.statusCode {
				t.Fatalf("expected %d, got %d", tc.statusCode, w.Code)
			}
		})
	}
}

func TestGetBooksController(t *testing.T) {
	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)

	if res["books"] == nil {
		t.Fatalf("expected books in response")
	}
}

func TestDeleteBookController(t *testing.T) {
	testCases := []struct {
		name       string
		body       string
		statusCode int
	}{
		{
			name:       "Existing Book",
			body:       `{"id":1}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Non-Existing Book",
			body:       `{"id":99}`,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "Invalid Body",
			body:       `{"id":"one"}`,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/delete", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.statusCode {
				t.Fatalf("expected %d, got %d", tc.statusCode, w.Code)
			}
		})
	}
}

func TestUpdateBookController(t *testing.T) {
	testCases := []struct {
		name       string
		body       string
		statusCode int
	}{
		{
			name:       "Valid Update",
			body:       `{"id":2,"title":"Updated Title","author":"Jane","genre":"Fiction","rating":4.8}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Non-Existing Book",
			body:       `{"id":99,"title":"Updated Title","author":"Jane","genre":"Fiction","rating":4.8}`,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "Invalid Body",
			body:       `{"id":2,"title":123,"author":"John","genre":"Tech","rating":4.5}`,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/update", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.statusCode {
				t.Fatalf("expected %d, got %d", tc.statusCode, w.Code)
			}
		})
	}
}
