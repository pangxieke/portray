package main

import (
	"fmt"
	"net/http"

	"github.com/pangxieke/portray/config"
	"github.com/pangxieke/portray/middle"
	"github.com/pangxieke/portray/model"
	"github.com/pangxieke/portray/router"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := model.Init(); err != nil {
		panic(err)
	}

	port := fmt.Sprintf("0.0.0.0:%d", config.Server.Port)
	fmt.Println("server start listen", port)
	err := http.ListenAndServe(port, middle.Handler(router.Router()))
	panic(err)
}
