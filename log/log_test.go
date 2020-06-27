package log_test

import (
	"github.com/pangxieke/portray/log"
	"testing"
)

func TestInit(t *testing.T) {
	log.Info("console log")

	err := log.Init("../error.log")
	if err != nil {
		t.Fatal(err)
	}
	log.Info("file log")
}
