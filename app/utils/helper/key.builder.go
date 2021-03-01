package helper

// KeyBuilder ...
func KeyBuilder(value ...string) string {
	key := ""

	for i, v := range value {
		if i == 0 {
			key += v
		} else {
			key += ":" + v
		}
	}

	return key
}
