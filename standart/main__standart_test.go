package standart

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResolver(t *testing.T) {
	req, _ := http.NewRequest("GET", "/resolver", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(resolverHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler return code: %v", rr.Code)
	}
}
