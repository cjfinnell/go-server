package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rr := httptest.NewRecorder()
	handleHello(rr, req)
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	rawBody, err := io.ReadAll(rr.Result().Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello!\n", string(rawBody))
}
