package shareddb

import (
	"time"

	"github.com/lib/pq"
)

type (
	BaseMachine struct {
		Base
		BaseOrg
		Name             string         `gorm:"column:name;unique" json:"name"`                //
		APIKey           string         `gorm:"column:api_key" json:"api_key"`                 // Машинаас сервэр-лүү хандалт хийх түлхүүр
		Status           string         `gorm:"column:status" json:"status"`                   // Машины төлөв 1-active 2-inactive 3-error
		POSTerminalID    string         `gorm:"column:pos_terminal_id" json:"pos_terminal_id"` //
		KeyNumber        string         `gorm:"column:key_number" json:"key_number"`           //
		WhiteKeyCount    uint           `gorm:"column:white_key_count" json:"white_key_count"` //
		BlackKeyCount    uint           `gorm:"column:black_key_count" json:"black_key_count"` //
		StaffID          uint           `gorm:"column:staff_id" json:"staff_id"`               // Байршил
		Staff            *Staff         `gorm:"foreignKey:StaffID" json:"staff,omitempty"`     //
		PossiblePayments pq.StringArray `gorm:"type:text[];column:possible_payments" json:"possible_payments,omitempty" swaggertype:"array,string"`
	}

	BaseTrip struct {
		Base
		BaseOrg
		DriverID uint   `gorm:"column:driver_id" json:"driver_id"`           // жолооч
		Driver   *Staff `gorm:"foreignKey:DriverID" json:"driver,omitempty"` //
	}

	BaseTripItem struct {
		Base
		BaseOrg
		NewSize     float64 `gorm:"column:new_size" json:"new_size"`         // 1- Хүлээгдэж байгаа, 2- Хийгдсэн //
		FillSize    float64 `gorm:"column:fill_size" json:"fill_size"`       // Дүүргэсэн хэмжээ
		ReturnSize  float64 `gorm:"column:return_size" json:"return_size"`   // Гаргасан хэмжээ
		InvoiceUUID string  `gorm:"column:invoice_uuid" json:"invoice_uuid"` // борлуулалт
		SellSize    float64 `gorm:"column:sell_size" json:"sell_size"`       // Борлуулалт гаргасан хэмжээ
		TripID      uint    `gorm:"column:trip_id" json:"trip_id"`           // харяалагдах аялал
		MachineID   uint    `gorm:"column:machine_id" json:"machine_id"`     // харяалагдах аялал
		SlotNumber  string  `gorm:"column:slot_number" json:"slot_number"`   // тухайн үеийн слотын дугаар
	}

	Location struct {
		Base
		BaseOrg
		Name               string            `gorm:"column:name;not null;unique" json:"name"`         //
		Image              string            `gorm:"column:image;" json:"image"`                      //
		Description        string            `gorm:"column:description" json:"description"`           //  Тайлбар
		AimagHot           string            `gorm:"column:aimag_hot" json:"aimag_hot"`               // Аймаг хот
		SumDuureg          string            `gorm:"column:sum_duureg" json:"sum_duureg"`             //
		BagHoroo           string            `gorm:"column:bag_horoo" json:"bag_horoo"`               // Баг хороо
		Latitude           float64           `gorm:"column:latitude" json:"latitude"`                 // Өргөрөг
		Longitude          float64           `gorm:"column:longitude" json:"longitude"`               // Уртраг
		ContactName        string            `gorm:"column:contact_name" json:"contact_name"`         // Байршилтай холбоо барих нэр
		ContactPhone       string            `gorm:"column:contact_phone" json:"contact_phone"`       // Байршилтай холбоо барих утас
		LastFilledDate     time.Time         `gorm:"column:last_filled_date" json:"last_filled_date"` // Сүүлд очсон огноо
		LocationCategoryID uint              `gorm:"column:location_category_id" json:"location_category_id"`
		LocationCategory   *LocationCategory `gorm:"foreignKey:LocationCategoryID" json:"location_category,omitempty"` //
	}

	LocationCategory struct {
		Base
		BaseOrg
		Name string `gorm:"column:name" json:"name"`
	}
)
