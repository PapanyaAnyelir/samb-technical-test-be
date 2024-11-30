package models

import (

	"log"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/samb_db_warehouse?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
	database.AutoMigrate(&Supplier{})
	database.AutoMigrate(&Customer{})
	database.AutoMigrate(&Warehouse{})
	database.AutoMigrate(&PenerimaanBarangHeader{})
	database.AutoMigrate(&PenerimaanBarangDetail{})
	database.AutoMigrate(&PengeluaranBarangHeader{})
	database.AutoMigrate(&PengeluaranBarangDetail{})

	DB = database

	log.Println("Menjalankan seeder...")
	if err := SeedProducts(); err != nil {
		log.Printf("Gagal menjalankan seeder Product: %v", err)
	}

	if err := SeedCustomers(); err != nil {
		log.Printf("Gagal menjalankan seeder Customer: %v", err)
	}

	if err := SeedWarehouses(); err != nil {
		log.Printf("Gagal menjalankan seeder Warehouse: %v", err)
	}

	if err := SeedSuppliers(); err != nil {
		log.Printf("Gagal menjalankan seeder Supplier: %v", err)
	}

	log.Println("Semua seeder berhasil dijalankan")
}
