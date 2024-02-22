package sharedutils

func GetUniqueStringsExclude(input []string, exclude string) []string {
	uniqueStrings := make(map[string]struct{})

	for _, s := range input {
		if s != exclude || s != "" {
			uniqueStrings[s] = struct{}{}
		}
	}

	result := make([]string, 0, len(uniqueStrings))
	for s := range uniqueStrings {
		result = append(result, s)
	}

	return result
}
