package shareddb

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Base struct {
		ID        uint           `gorm:"primary_key;autoIncrement:true" json:"id"`
		CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	}
	BaseOrg struct {
		OrgID uint          `gorm:"column:org_id" json:"org_id"`
		Org   *Organization `gorm:"foreignKey:OrgID" json:"org,omitempty"`
	}

	BaseLang struct {
		ReferralID uint   `gorm:"column:referral_id" json:"referral_id"`
		Lang       string `gorm:"column:lang" json:"lang"`
	}

	PaymentType struct {
		Base
		Name             string         `gorm:"column:name" json:"name"`               // Нэр
		TypeString       string         `gorm:"column:type_string" json:"type_string"` // string
		ConfigEnvStrings pq.StringArray `gorm:"type:text[];column:config_env_strings" json:"config_env_strings,omitempty" swaggertype:"array,string"`
		ColorCode        string         `gorm:"column:color_code" json:"color_code"` // Өнгө
	}

	PaymentConfig struct {
		Base
		OrgID                  uint          `gorm:"column:org_id" json:"org_id"`
		Organization           *Organization `gorm:"foreignKey:OrgID" json:"organization,omitempty"`
		PaymentTypeID          uint          `gorm:"column:payment_type_id" json:"payment_type_id"`
		PaymentType            *PaymentType  `gorm:"foreignKey:PaymentTypeID" json:"payment_type,omitempty"`
		IsActive               bool          `gorm:"column:is_active" json:"is_active"`
		DiscountPricePrecent   float64       `gorm:"column:discount_price_precent" json:"discount_percent"` // Хямдралын хувь
		DiscountEbarimtPrecent float64       `gorm:"column:discount_ebarimt_precent" json:"discount_ebarimt_precent"`
		UUID                   string        `gorm:"column:uuid" json:"uuid"`
		Configs                []Config      `gorm:"foreignKey:RefUUID;references:UUID" json:"configs,omitempty"`
	}

	Lang struct {
		Base
		Name string `gorm:"column:name" json:"name"`
		Code string `gorm:"column:code;unique" json:"code"`
	}

	Address struct {
		Base
		ParentID     uint   `gorm:"column:parent_id" json:"parent_id"`
		RegionName   string `gorm:"column:region_name" json:"region_name"`
		RegionTypeID int    `gorm:"column:region_type_id" json:"region_type_id"`
		OrderField   int    `gorm:"column:order_field;not null;autoIncrement" json:"order_field"`
	}
)
