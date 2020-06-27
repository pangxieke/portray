package view

import (
	"github.com/pangxieke/portray/model"
)

func NewDevice(a model.Device) (res interface{}) {
	v := map[string]interface{}{
		"id":         a.ID,
		"sn":         a.SN,
		"mac":        a.Mac,
		"created_at": a.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": a.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return v
}
