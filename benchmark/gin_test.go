package benchmark

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/hello/:name", func(c *gin.Context) {
		c.String(200, "ok")
	})

	req := httptest.NewRequest("GET", "/hello/rishabh", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r.ServeHTTP(w, req)
	}
}
