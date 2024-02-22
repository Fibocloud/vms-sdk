package sharedutils

import (
	"fmt"
	"strings"
)

func FormatFullName(firstName, lastName string) string {
	// Extract the first letter from the first name
	lastInitial := string(lastName[0])

	// Concatenate the initial with the last name
	formattedName := fmt.Sprintf("%s. %s", lastInitial, firstName)

	// Capitalize the first letter of each word
	formattedName = strings.Title(formattedName)

	return formattedName
}
