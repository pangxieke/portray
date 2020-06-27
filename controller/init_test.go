package controller_test

import (
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
	portrayRou "github.com/pangxieke/portray/router"
	"github.com/pangxieke/portray/test"
)

var (
	router *httprouter.Router
)

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	router = portrayRou.Router()
	test.Init()
}
