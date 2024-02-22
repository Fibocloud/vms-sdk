package sharedutils

import (
	"fmt"
	"strings"
)

func GeneratePermissionCode(method, path string) string {
	if strings.Contains(path, "/api/v1/") {
		path = path[7:]
	}
	if strings.Contains(path, "uabz/") {
		path = path[5:]
	}
	return fmt.Sprintf("%s:%s", strings.ToLower(method), strings.ToLower(path))
}
