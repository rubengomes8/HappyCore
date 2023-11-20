package collections

// ListValues returns the map values as a list.
func ListValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ToMap given a list, returns the list content as a map.
func ToMap[K comparable, V any](list []V, getKeyValue func(V) K) map[K]V {
	newMap := make(map[K]V)
	for _, current := range list {
		newMap[getKeyValue(current)] = current
	}
	return newMap
}

// MapKeys returns an array with the map keys.
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
