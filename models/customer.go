package models

type Customer struct {
	CustomerPK           int64  `gorm:"primaryKey" json:"customer_pk"`
	CustomerName 			string `gorm:"type:varchar(300)" json:"customer_name"`
}

func SeedCustomers() error {
	customers := []Customer{
		{CustomerName: "Customer A"},
		{CustomerName: "Customer B"},
		{CustomerName: "Customer C"},
	}

	for _, customer := range customers {
		if err := DB.FirstOrCreate(&Customer{}, customer).Error; err != nil {
			return err
		}
	}

	return nil
}
