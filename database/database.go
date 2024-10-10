package database

import (
	"log"

	"github.com/Nivas-Mekala/IKEA_Inventory_Application/internal/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectToDatabase() {

	dsn := "root:password@tcp(127.0.0.1:3306)/ikea_app?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error Connecting database %s \n", err.Error())
	}
	log.Println("Connected to Database")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations!!")
	db.AutoMigrate(&models.Product{}, &models.Inventory{})
	//db.Table("Inventory").AutoMigrate(&models.Inventories{})
	//db.Table("Products").AutoMigrate(&models.Product{})

	Database = DBInstance{
		Db: db,
	}
}
