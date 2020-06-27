package config_test

import (
	"testing"

	"github.com/pangxieke/portray/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	err := config.Init("../")
	assert.Nil(err)

	port := config.Server.Port
	assert.NotEqual(0, port)
}
