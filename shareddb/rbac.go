package shareddb

import (
	"strings"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	// Permission [ Эрхийн тохиргоо ]
	Permission struct {
		Base
		Code           string `gorm:"column:code;unique" json:"code"`
		Name           string `gorm:"column:name" json:"name"`
		Path           string `gorm:"column:path;not null" json:"path"` // API  зам
		Method         string `gorm:"column:method" json:"method"`      // API  зам
		IsVisibleToOrg bool   `gorm:"column:is_visible_to_org;default:false" json:"is_visible_to_org"`
	}

	Role struct {
		Base
		BaseOrg
		Name          string        `gorm:"column:name;not null" json:"name"`
		PermissionIDs pq.Int64Array `gorm:"type:integer[];column:permission_ids" json:"permission_ids,omitempty" swaggertype:"array,number" example:"1,2"`
		GroupIDs      pq.Int64Array `gorm:"type:integer[];column:group_ids" json:"group_ids,omitempty" swaggertype:"array,number" example:"1,2"`
	}

	Group struct {
		Base
		Key            string        `gorm:"column:key;not null" json:"key"`
		Name           string        `gorm:"column:name;not null" json:"name"`
		ParentID       uint          `gorm:"column:parent_id" json:"parent_id,omitempty"`
		IsParent       bool          `gorm:"column:is_parent" json:"is_parent"`
		PermissionIDs  pq.Int64Array `gorm:"type:integer[];column:permission_ids" json:"permission_ids,omitempty" swaggertype:"array,number" example:"1,2"`
		IsVisibleToOrg bool          `gorm:"column:is_visible_to_org;default:false" json:"is_visible_to_org"`
	}
)

func (instance *Role) IsGranted(fullPath, method string, db *gorm.DB) bool {
	if instance == nil {
		return false
	}
	if strings.ToLower(instance.Name) == "admin" {
		return true
	}

	fullPath = fullPath[7:]
	permission := Permission{}
	if err := db.Where("path = ? and method = ?", fullPath, method).First(&permission).Error; err != nil {
		return true
	}
	for _, v := range instance.PermissionIDs {
		if v == int64(permission.ID) {
			return true
		}
	}

	return false
}
