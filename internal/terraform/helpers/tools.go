package helpers

// KeyInMap checks if a string key is in a map
func KeyInMap(key string, value map[string]any) bool {
	for k := range value {
		if k == key {
			return true
		}
	}
	return false
}

// ListContains checks if string is in list
func ListContains(list []string, value string) bool {
	for _, elem := range list {
		if elem == value {
			return true
		}
	}
	return false
}
