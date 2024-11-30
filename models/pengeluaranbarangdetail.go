package models

type PengeluaranBarangDetail struct {
	TrxOutDPK         int64 `gorm:"primaryKey" json:"trx_out_dpk"`
	TrxOutIDF         int64 `gorm:"not null" json:"trx_out_idf"`
	TrxOutDProductIdf int64 `gorm:"not null" json:"trx_out_d_product_idf"`
	TrxOutDQtyDus     int   `gorm:"not null" json:"trx_out_d_qty_dus"`
	TrxOutDQtyPcs     int   `gorm:"not null" json:"trx_out_d_qty_pcs"`
}
