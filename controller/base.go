package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Base struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Params         httprouter.Params
	RequestBody    string
}

func (b *Base) Init(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	b.Request, b.ResponseWriter, b.Params = r, rw, p
	if b.Request.ContentLength > 0 {
		buffer := make([]byte, b.Request.ContentLength)
		b.Request.Body.Read(buffer)
		b.RequestBody = string(buffer)
	}

	return nil
}

func (b *Base) Destroy() {
}

func (b *Base) Error(err error) {
	http.Error(b.ResponseWriter, err.Error(), http.StatusInternalServerError)
}
