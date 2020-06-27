package controller

import (
	"github.com/pangxieke/portray/log"
	"github.com/pangxieke/portray/model"
	"github.com/pangxieke/portray/view"
)

type Device struct {
	Application
}

func (this *Device) List() (err error) {
	data, err := model.FindDevList()
	if err != nil {
		log.Info("FindDevList error:", err)
	}
	return this.respondJson(data)
}

func (this *Device) Info() (err error) {
	id, err := this.getID("id")
	if err != nil {
		return
	}

	if id == 0 {
		return newNotFoundError("id", id)
	}

	data, err := model.FindDevById(uint(id))
	if err != nil {
		log.Info("FindDevById error:", err)
	}
	v := view.NewDevice(*data)
	return this.respondJson(v)
}
