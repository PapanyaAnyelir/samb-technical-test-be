package models

type PenerimaanBarangDetail struct {
	TrxInDPK         int64 `gorm:"primaryKey" json:"trx_in_dpk"`
	TrxInIDF         int64 `gorm:"not null" json:"trx_in_idf"`
	TrxInDProductIdf int64 `gorm:"not null" json:"trx_in_d_product_idf"`
	TrxInDQtyDus     int   `gorm:"not null" json:"trx_in_d_qty_dus"`
	TrxInDQtyPcs     int   `gorm:"not null" json:"trx_in_d_qty_pcs"`
}
