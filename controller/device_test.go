package controller_test

import (
	"encoding/json"
	"github.com/pangxieke/portray/model"
	"github.com/pangxieke/portray/test"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDevice_List(t *testing.T) {
	assert := assert.New(t)
	test.Prepare(t)

	//test.Init()
	requestBody := strings.NewReader(``)
	request, _ := http.NewRequest("GET", "/devices", requestBody)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(http.StatusOK, response.Code, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(err)

	var data []model.Device
	assert.Nil(json.Unmarshal(body, &data), string(body))

	assert.NotEmpty(data)
	expect := []model.Device{
		{ID: 1, SN: "TestSN00001", Mac: "85:20:96:CC:5D:95"},
		{ID: 2, SN: "TestSN00002", Mac: "85:20:96:CC:5D:96"},
	}
	for key, val := range data {
		assert.EqualValues(val.ID, expect[key].ID)
		assert.EqualValues(val.SN, expect[key].SN)
		assert.EqualValues(val.Mac, expect[key].Mac)
	}

}

func TestDevice_Info(t *testing.T) {
	assert := assert.New(t)
	test.Prepare(t)

	requestBody := strings.NewReader(``)
	request, _ := http.NewRequest("GET", "/device/1", requestBody)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(http.StatusOK, response.Code, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(err)

	var data model.Device
	assert.Nil(json.Unmarshal(body, &data), string(body))

	assert.EqualValues(1, data.ID)
	assert.Equal("TestSN00001", data.SN)
}
