package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	ch := NewCommonhandler()
	r.GET("/", ch.Root)
	r.GET("/ping", ch.Ping)
	r.GET("/health", ch.Health("OK"))
	return r
}
func Test_Root(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	r := setUpRouter()

	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "Hello World!", rec.Body.String())
}

func Test_Ping(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	r := setUpRouter()

	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, rec.Header().Get("Content-Type"), "application/json; charset=utf-8")

	var got map[string]any

	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))

	require.Equal(t, "pong", got["ping"])

	// Content-Type application/json; charset=utf-8
	// Content-Length 15
}
