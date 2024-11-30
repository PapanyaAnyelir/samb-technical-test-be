package suppliercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/PapanyaAnyelir/samb-technical-test-be/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var suppliers []models.Supplier

	models.DB.Find(&suppliers)
	c.JSON(http.StatusOK, gin.H{"suppliers": suppliers})

}

func Show(c *gin.Context) {
	var supplier models.Supplier
	id := c.Param("id")

	if err := models.DB.First(&supplier, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"supplier": supplier})
}

func Create(c *gin.Context) {

	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&supplier)
	c.JSON(http.StatusOK, gin.H{"supplier": supplier})
}

func Update(c *gin.Context) {
	var supplier models.Supplier
	id := c.Param("id")

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&supplier).Where("supplier_pk = ?", id).Updates(&supplier).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var supplier models.Supplier

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&supplier, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
