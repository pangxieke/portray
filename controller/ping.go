package controller

import (
	"errors"
)

type PingController struct {
	Base
}

func (this *PingController) Ping() (err error) {
	_, err = this.ResponseWriter.Write([]byte("pong \n"))
	return
}

func (this *PingController) Health() (err error) {
	err = errors.New("fail")
	return
}
