package models

import (
	"time"
)	

type PenerimaanBarangHeader struct {
	TrxInPK      int64     `gorm:"primaryKey" json:"trx_in_pk"`
	TrxInNo      string    `gorm:"type:varchar(100);not null" json:"trx_in_no"`
	WhsIdf       int64     `gorm:"not null" json:"whs_idf"`
	TrxInDate    time.Time `gorm:"not null" json:"trx_in_date"`
	TrxInSuppIdf int64     `gorm:"not null" json:"trx_in_supp_idf"`
	TrxInNotes   string    `gorm:"type:varchar(255)" json:"trx_in_notes"`
}
