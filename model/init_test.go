package model

import (
	"github.com/pangxieke/portray/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	err := config.Init("../")
	assert.Nil(err)

	err = Init()
	if err != nil {
		t.Fatal(err)
	}
}
