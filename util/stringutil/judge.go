package stringutil

// ContainsString a string in string list
func ContainsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}
