package benchmark

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func BenchmarkMux(b *testing.B) {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}).Methods("GET")

	req := httptest.NewRequest("GET", "/hello/rishabh", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r.ServeHTTP(w, req)
	}
}
