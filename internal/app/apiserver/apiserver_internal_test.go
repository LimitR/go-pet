package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIserver_HangleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
}
