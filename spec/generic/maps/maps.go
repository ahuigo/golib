package maps

func GetValue[K comparable, V any](m map[K]V, key K, defaultV V) V {
	if m == nil {
		return defaultV
	}
	if v, ok := m[key]; ok {
		return v
	}
	return defaultV
}
