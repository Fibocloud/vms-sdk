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
)
