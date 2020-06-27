package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	Base
	UUID string
}

func (a *Application) Init(rw http.ResponseWriter, r *http.Request, p httprouter.Params) (err error) {
	if err = a.Base.Init(rw, r, p); err != nil {
		return
	}

	//request uuid
	id, _ := uuid.NewUUID()
	a.UUID = fmt.Sprintf("%s", id)

	return
}

//implement Error method of controller interface
func (a *Application) Error(err error) {
	r := map[string]interface{}{
		"message":   err.Error(),
		"requestId": a.UUID,
		"code":      0,
	}
	//err implement Error struct
	if e, ok := err.(Error); ok {
		if e.Msg != "" {
			r["message"] = e.Msg
		}
		if e.Code != nil {
			r["code"] = e.Code
		}

	}

	a.respondJson(r, errorHttpStatus(err))
}

func (a *Application) respondJson(response interface{}, statusCode ...int) error {
	if len(statusCode) > 0 {
		a.ResponseWriter.WriteHeader(statusCode[0])
	}

	j, err := json.Marshal(response)
	if err != nil {
		return err
	}
	a.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	a.ResponseWriter.Write(j)
	return nil
}

func (a *Application) getID(name string) (res int, err error) {
	id := a.getPathParam("id")
	if id == "" {
		return 0, newInvalidParameterError("id", "")
	}
	res, _ = strconv.Atoi(id)
	return
}

// page offset limit
func (a *Application) getLimitAndOffset() (limit, offset int, err error) {
	offset, err = a.getQueryInt("offset", 0)
	if err != nil {
		return
	}
	limit, err = a.getQueryInt("limit", 24)
	if err != nil {
		return
	}
	return
}

func (a *Application) extractParams(params interface{}) (err error) {
	err = json.Unmarshal([]byte(a.RequestBody), params)
	if err != nil {
		log.Printf("invalid body format, err = %+v\n", err)
		return newInvalidInputErrorf("invalid body format")
	}
	return
}

func (b *Application) getPathParam(name string) (res string) {
	for _, p := range b.Params {
		if p.Key == name {
			return p.Value
		}
	}
	return
}

func (a *Application) getQuery(name string) (res string, err error) {
	values, ok := a.Request.URL.Query()[name]
	if !ok || len(values[0]) < 1 {
		err = newInvalidParameterError(name, "missing")
		return
	}
	res = values[0]
	return
}

func (a *Application) getQueryInt(name string, defaultValue ...int) (res int, err error) {
	value, err := a.getQuery(name)
	if err != nil {
		return
	}
	if value == "" {
		if len(defaultValue) > 1 {
			return defaultValue[0], nil
		}
		return 0, newInvalidParameterError(name, "")
	}
	res, err = strconv.Atoi(value)
	if err != nil {
		return 0, newInvalidParameterError(name, err.Error())
	}
	return
}
