// models/setup.go
package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=Habibul77 dbname=restexpress port=5432 sslmode=disable timezone=Asia/Jakarta"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  // change the database provider if necessary

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Post{})  // register Post model

    DB = database
}