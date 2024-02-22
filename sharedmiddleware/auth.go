package sharedmiddleware

import (
	"fmt"

	"github.com/Fibocloud/vms-sdk/constants"
	"github.com/Fibocloud/vms-sdk/shareddb"
	"github.com/Fibocloud/vms-sdk/sharedutils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const authKey string = "auth"
const orgKey string = "org"

// Auth - middleware from token
func Auth(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		requiredToken, ok := headers["Authorization"]
		if !ok || len(requiredToken) == 0 || len(requiredToken[0]) < 8 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Таны нэвтрэх хэрэгтэй",
				"body":    nil,
			})
		}

		orgHeaderID := headers["Org-Id"]

		claims, err := sharedutils.ExtractJWTString(requiredToken[0][7:])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Таны нэвтрэх хугацаа дууссан байна",
				"body":    nil,
			})
		}
		var user shareddb.Staff
		if err := db.Preload("Org").Preload("Role").First(&user, claims.ID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
				"body":    nil,
			})
		}

		org := shareddb.Organization{}
		orgID := ""
		if len(orgHeaderID) == 0 || orgHeaderID[0] == "" {
			orgID = fmt.Sprintf("%d", user.OrgID)
		}
		db.First(&org, orgID)

		if user.OrgID != org.ID && user.Org.Level != constants.OrgLevelTop {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Хандах эрх хүрэхгүй байна.",
				"body":    nil,
			})
		}

		if user.Role.IsGranted(c.OriginalURL(), c.Method(), db) {
			c.Locals(orgKey, &org)
			c.Locals(authKey, &user)
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Хандах эрх хүрэхгүй байна.",
			"body":    nil,
		})
	}
}

// GetAuth - Get auth user from context
func GetAuth(c *fiber.Ctx) *shareddb.Staff {
	return c.Locals(authKey).(*shareddb.Staff)
}

// GetOrg
func GetOrg(c *fiber.Ctx) *shareddb.Organization {
	return c.Locals(orgKey).(*shareddb.Organization)
}

// GetAuthOptional
func GetAuthOptional(c *fiber.Ctx, db *gorm.DB) *shareddb.Staff {
	headers := c.GetReqHeaders()
	requiredToken := headers["Authorization"]
	if len(requiredToken) == 0 || len(requiredToken[0]) < 8 {
		return nil
	}
	claims, err := sharedutils.ExtractJWTString(requiredToken[0][7:])
	if err != nil {
		return nil
	}
	var user shareddb.Staff
	if err := db.First(&user, claims.ID).Error; err != nil {
		return nil
	}
	return &user
}
