package collections

// MapList transforms a list of type V into a list of type T.
func MapList[V, T any](originalList []V, transform func(V) T) []T {
	newList := make([]T, 0, len(originalList))
	for _, current := range originalList {
		newList = append(newList, transform(current))
	}
	return newList
}

// DistinctV0 returns a list of distinct elements based on a given field.
func DistinctV0[V any](list []V, field func(V) any) []V {
	uniqueMap := make(map[interface{}]V)
	for _, current := range list {
		uniqueMap[field(current)] = current
	}

	newList := make([]V, 0, len(uniqueMap))
	for _, current := range uniqueMap {
		newList = append(newList, current)
	}

	return newList
}

// DistinctV1 returns a list of distinct elements based on a given field.
func DistinctV1[V any, K comparable](list []V, getKeyFunc func(V) K) []V {
	mp := ToMap(list, getKeyFunc)
	return ListValues(mp)
}

// Filter returns a new list with the filtered values.
func Filter[V any](list []V, filterFunc func(V) bool) []V {
	newList := make([]V, 0, len(list))
	for _, current := range list {
		if filterFunc(current) {
			newList = append(newList, current)
		}
	}
	return newList
}

// Contains returns true if the element is on the list.
func Contains[V comparable](slice []V, elem V) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
