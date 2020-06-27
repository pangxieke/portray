package model

func Migrate() error {
	return db.AutoMigrate(&Device{}).Error
}
