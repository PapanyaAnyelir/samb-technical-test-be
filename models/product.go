package models

type Product struct {
	ProductPK          int64  `gorm:"primaryKey" json:"product_pk"`
	ProductName string `gorm:"type:varchar(300)" json:"product_name"`
}

func SeedProducts() error {
	products := []Product{
		{ProductName: "Product A"},
		{ProductName: "Product B"},
		{ProductName: "Product C"},
	}

	for _, product := range products {
		if err := DB.Create(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
