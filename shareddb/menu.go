package shareddb

type (
	BaseProduct struct {
		Base
		BaseOrg
		Name         string  `gorm:"->;type:varchar(1)" json:"name"`
		VeritechCode string  `gorm:"column:veritech_code" json:"veritech_code"` //
		ImagePath    string  `gorm:"column:image_path" json:"image_path"`       // Зураг         //
		Measurement  string  `gorm:"column:measurement" json:"measurement"`     // Хэмжих нэгж
		Price        float64 `gorm:"column:price" json:"price"`                 // Vending машин дээр харагдах үнэ (нийт)
		IsActive     bool    `gorm:"column:is_active" json:"is_active"`         // Бүтээгдэхүүн идэвхтэй эсэх
	}
)
