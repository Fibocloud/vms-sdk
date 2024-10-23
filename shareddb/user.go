package shareddb

import "time"

type (
	Staff struct {
		Base
		BaseOrg
		Code          string    `gorm:"column:code;unique;not null" json:"code"`
		Email         string    `gorm:"column:email;unique;not null" json:"email"`
		FirstName     string    `gorm:"column:first_name;not null" json:"first_name"`
		LastName      string    `gorm:"column:last_name" json:"last_name"`
		Phone         string    `gorm:"column:phone;size:12" json:"phone"`
		Password      string    `gorm:"column:password;not null" json:"-"`
		LastLoginDate time.Time `gorm:"column:last_login_date" json:"last_login_date"`
		IsActive      bool      `gorm:"column:is_active;not null;default:true" json:"is_active"`
		HaveWebAccess bool      `gorm:"column:have_web_access" json:"have_web_access"`
		DeviceID      string    `gorm:"column:device_id" json:"device_id"`
		RoleID        uint      `gorm:"column:role_id" json:"role_id"`           //
		Role          *Role     `gorm:"foreignKey:RoleID" json:"role,omitempty"` //
	}

	Organization struct {
		Base
		Name          string `gorm:"column:name;not null" json:"name"`      // Нэр
		Description   string `gorm:"column:description" json:"description"` // Тайлбар
		Level         string `gorm:"column:level" json:"level"`             //
		FacebookName  string `gorm:"column:facebook_name" json:"facebook_name"`
		InstagramName string `gorm:"column:instagram_name" json:"instagram_name"`
		Email         string `gorm:"column:email" json:"email"`
		ContactName   string `gorm:"column:contact_name" json:"contact_name"`
		Phone         string `gorm:"column:phone" json:"phone"`
	}
)
