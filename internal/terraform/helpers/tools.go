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
