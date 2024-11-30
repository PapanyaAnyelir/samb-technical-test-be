package customercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/PapanyaAnyelir/samb-technical-test-be/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var customer models.Customer

func Index(c *gin.Context) {
	var customers []models.Customer

	models.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"customers": customers})

}

func Show(c *gin.Context) {
	id := c.Param("id")

	if err := models.DB.First(&customer, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

func Create(c *gin.Context) {

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&customer).Where("customer_pk = ?", id).Updates(&customer).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate Customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&customer, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
