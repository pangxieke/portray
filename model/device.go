package model

import "time"

type Device struct {
	ID        uint   `gorm:"primary_key"`
	SN        string `gorm:"column:sn"`
	Mac       string `gorm:"column:mac"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Device) TableName() string {
	return "device"
}

func FindDevById(id uint) (item *Device, err error) {
	var u Device
	s := db.Where("id = ?", id).First(&u)
	if s.RecordNotFound() {
		return
	}
	if err = s.Error; err != nil {
		return
	}
	item = &u
	return
}

func FindDevList() (item *[]Device, err error) {
	var u []Device
	s := db.
		Where("created_at < ?", time.Now()).
		Find(&u)
	if s.RecordNotFound() {
		return
	}
	if err = s.Error; err != nil {
		return
	}
	item = &u
	return
}

func (m *Device) Save() (err error) {
	m.UpdatedAt = time.Now()
	return db.Save(m).Error
}
