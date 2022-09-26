package database

import (
	"fmt"

	"github.com/normos/qrresume/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbConn() *gorm.DB {
	dsn := "host=localhost user=postgres password=mit123 dbname=qr port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully connected")
	}
	return db
}

func DbMigration(d *gorm.DB) {
	d.AutoMigrate(&models.User{})
}

func CreateUser(db *gorm.DB, u *models.Usercreateac) {
	db.Create(&models.User{Name: u.Name, EmailId: u.EmailId, Password: u.Password})

}
