package models

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }


    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbCharset := os.Getenv("DB_CHARSET")
    dbParseTime := os.Getenv("DB_PARSE_TIME")
    dbLoc := os.Getenv("DB_LOC")


    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
        dbUser, dbPass, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc)


    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    database.AutoMigrate(&Product{}, &Supplier{}, &Customer{}, &Warehouse{},
        &PenerimaanBarangHeader{}, &PenerimaanBarangDetail{},
        &PengeluaranBarangHeader{}, &PengeluaranBarangDetail{})

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
