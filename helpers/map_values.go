package helpers

func MapValues[K comparable, V any](m map[K]V) []V {
	result := []V{}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
