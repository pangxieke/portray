package controller_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingController_Ping(t *testing.T) {
	assert := assert.New(t)

	requestBody := strings.NewReader(``)
	request, _ := http.NewRequest("GET", "/ping", requestBody)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(http.StatusOK, response.Code, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(err)
	assert.Equal("pong \n", string(body))
}
