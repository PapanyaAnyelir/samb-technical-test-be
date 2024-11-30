package warehousecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/PapanyaAnyelir/samb-technical-test-be/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var warehouses []models.Warehouse

	models.DB.Find(&warehouses)
	c.JSON(http.StatusOK, gin.H{"warehouses": warehouses})

}

func Show(c *gin.Context) {
	var warehouse models.Warehouse
	id := c.Param("id")

	if err := models.DB.First(&warehouse, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"warehouse": warehouse})
}

func Create(c *gin.Context) {

	var warehouse models.Warehouse

	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&warehouse)
	c.JSON(http.StatusOK, gin.H{"warehouse": warehouse})
}

func Update(c *gin.Context) {
	var warehouse models.Warehouse
	id := c.Param("id")

	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&warehouse).Where("whs_pk = ?", id).Updates(&warehouse).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate warehouse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var warehouse models.Warehouse

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&warehouse, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus warehouse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
