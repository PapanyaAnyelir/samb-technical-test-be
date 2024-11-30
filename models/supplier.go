package models

type Supplier struct {
	SupplierPK          int64  `gorm:"primaryKey" json:"supplier_pk"`
	SupplierName string `gorm:"type:varchar(300)" json:"supplier_name"`
}

func SeedSuppliers() error {
	suppliers := []Supplier{
		{SupplierName: "Supplier A"},
		{SupplierName: "Supplier B"},
		{SupplierName: "Supplier C"},
	}

	for _, supplier := range suppliers {
		if err := DB.FirstOrCreate(&Supplier{}, supplier).Error; err != nil {
			return err
		}
	}

	return nil
}


