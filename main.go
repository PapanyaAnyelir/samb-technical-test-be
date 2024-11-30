package main

import (
    "github.com/PapanyaAnyelir/samb-technical-test-be/controllers/productcontroller"
    "github.com/PapanyaAnyelir/samb-technical-test-be/controllers/customercontroller"
    "github.com/PapanyaAnyelir/samb-technical-test-be/controllers/suppliercontroller"
    "github.com/PapanyaAnyelir/samb-technical-test-be/controllers/warehousecontroller"
    "github.com/PapanyaAnyelir/samb-technical-test-be/controllers/transactioncontroller"
    "github.com/PapanyaAnyelir/samb-technical-test-be/models"

    "github.com/gin-contrib/cors" 
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    models.ConnectDatabase()

    // Middleware CORS
    r.Use(cors.New(cors.Config{
		  AllowOrigins: []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Routes
    r.GET("/api/products", productcontroller.Index)
    r.GET("/api/product/:id", productcontroller.Show)
    r.POST("/api/product", productcontroller.Create)
    r.PUT("/api/product/:id", productcontroller.Update)
    r.DELETE("/api/product", productcontroller.Delete)

    r.GET("/api/customers", customercontroller.Index)
    r.GET("/api/customer/:id", customercontroller.Show)
    r.POST("/api/customer", customercontroller.Create)
    r.PUT("/api/customer/:id", customercontroller.Update)
    r.DELETE("/api/customer", customercontroller.Delete)

    r.GET("/api/suppliers", suppliercontroller.Index)
    r.GET("/api/supplier/:id", suppliercontroller.Show)
    r.POST("/api/supplier", suppliercontroller.Create)
    r.PUT("/api/supplier/:id", suppliercontroller.Update)
    r.DELETE("/api/supplier", suppliercontroller.Delete)

    r.GET("/api/warehouses", warehousecontroller.Index)
    r.GET("/api/warehouse/:id", warehousecontroller.Show)
    r.POST("/api/warehouse", warehousecontroller.Create)
    r.PUT("/api/warehouse/:id", warehousecontroller.Update)
    r.DELETE("/api/warehouse", warehousecontroller.Delete)

    r.GET("/api/laporan-stok", transactioncontroller.LaporanStok)
    r.POST("/api/inbound", transactioncontroller.Inbound)
    r.POST("/api/outbound", transactioncontroller.Outbound)

    r.Run()
}
