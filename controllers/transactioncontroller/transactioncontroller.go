package transactioncontroller

import (
	"net/http"
	"time"

	"github.com/PapanyaAnyelir/samb-technical-test-be/models"
	"github.com/PapanyaAnyelir/samb-technical-test-be/utils"
	"github.com/gin-gonic/gin"

)

func Inbound(c *gin.Context) {
	var payload struct {
		Header  models.PenerimaanBarangHeader   `json:"header"`
		Details []models.PenerimaanBarangDetail `json:"details"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tx := models.DB.Begin()

	payload.Header.TrxInNo = utils.GenerateInboundNumber()
	payload.Header.TrxInDate = time.Now()

	if err := tx.Create(&payload.Header).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan header"})
		return
	}

	for _, detail := range payload.Details {
		detail.TrxInIDF = payload.Header.TrxInPK
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan detail"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Penerimaan barang berhasil disimpan", "trx_in_no": payload.Header.TrxInNo})
}


func Outbound(c *gin.Context) {
	var payload struct {
		Header  models.PengeluaranBarangHeader   `json:"header"`
		Details []models.PengeluaranBarangDetail `json:"details"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tx := models.DB.Begin()

	payload.Header.TrxOutNo = utils.GenerateOutboundNumber()
	payload.Header.TrxOutDate = time.Now()

	if err := tx.Create(&payload.Header).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan header pengeluaran"})
		return
	}

	for _, detail := range payload.Details {
		detail.TrxOutIDF = payload.Header.TrxOutPK
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan detail pengeluaran"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Pengeluaran barang berhasil disimpan", "trx_out_no": payload.Header.TrxOutNo})
}


func LaporanStok(c *gin.Context) {
	type Stok struct {
		WarehouseName string `json:"warehouse_name"`
		ProductName   string `json:"product_name"`
		QtyDus        int    `json:"qty_dus"`
		QtyPcs        int    `json:"qty_pcs"`
	}

	var stok []Stok

	query := `
		SELECT 
			w.whs_name AS warehouse_name,
			p.product_name AS product_name,
			COALESCE(SUM(in_detail.trx_in_d_qty_dus), 0) - COALESCE(SUM(out_detail.trx_out_d_qty_dus), 0) AS qty_dus,
			COALESCE(SUM(in_detail.trx_in_d_qty_pcs), 0) - COALESCE(SUM(out_detail.trx_out_d_qty_pcs), 0) AS qty_pcs
		FROM 
			warehouses w
		CROSS JOIN 
			products p
		LEFT JOIN 
			penerimaan_barang_details in_detail ON p.product_pk = in_detail.trx_in_d_product_idf
		LEFT JOIN 
			penerimaan_barang_headers in_header ON in_detail.trx_in_id_f = in_header.trx_in_pk AND in_header.whs_idf = w.whs_pk
		LEFT JOIN 
			pengeluaran_barang_details out_detail ON p.product_pk = out_detail.trx_out_d_product_idf
		LEFT JOIN 
			pengeluaran_barang_headers out_header ON out_detail.trx_out_id_f = out_header.trx_out_pk AND out_header.whs_idf = w.whs_pk
		GROUP BY 
			w.whs_name, p.product_name`

	if err := models.DB.Raw(query).Scan(&stok).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil laporan stok"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stok": stok})
}

