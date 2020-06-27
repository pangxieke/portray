package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pangxieke/portray/controller"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/ping", controller.Action((*controller.PingController).Ping))
	router.GET("/health", controller.Action((*controller.PingController).Health))
	router.GET("/devices", controller.Action((*controller.Device).List))
	router.GET("/device/:id", controller.Action((*controller.Device).Info))

	return router

}
