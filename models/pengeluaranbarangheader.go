package models

import (
	"time"
)

type PengeluaranBarangHeader struct {
	TrxOutPK       int64                     `gorm:"primaryKey" json:"trx_out_pk"`
	TrxOutNo       string                    `json:"trx_out_no"`
	WhsIdf         int                       `json:"whs_idf"`
	TrxOutDate     time.Time                 `json:"trx_out_date"`
	TrxOutSuppIdf  int                       `json:"trx_out_supp_idf"`
	TrxOutNotes    string                    `json:"trx_out_notes"`
	Details        []PengeluaranBarangDetail `json:"details" gorm:"-"`
}
