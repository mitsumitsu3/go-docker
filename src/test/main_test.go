package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"www/src/router"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	rt := router.Init()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	rt.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())
}
