package model

func migration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Address{})
	DB.AutoMigrate(&Tel{})
	DB.AutoMigrate(&Language{})
}
