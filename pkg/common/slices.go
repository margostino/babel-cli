package common

func AppendIfNotExists(slice []string, element string) []string {
	for _, existingElement := range slice {
		if existingElement == element {
			return slice
		}
	}
	return append(slice, element)
}
