package models

type Warehouse struct {
	WhsPK          int64  `gorm:"primaryKey" json:"warehouse_pk"`
	WhsName string `gorm:"type:varchar(300)" json:"warehouse_name"`
}


func SeedWarehouses() error {
	warehouses := []Warehouse{
		{WhsName: "Warehouse A"},
		{WhsName: "Warehouse B"},
		{WhsName: "Warehouse C"},
	}

	for _, warehouse := range warehouses {
		if err := DB.FirstOrCreate(&Warehouse{}, warehouse).Error; err != nil {
			return err
		}
	}

	return nil
}

