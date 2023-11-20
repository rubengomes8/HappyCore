package collections

// FirstIndex returns the index of the first element that checks.
func FirstIndex[V any](list []V, checkFunc func(V) bool) int {
	for i, cur := range list {
		if checkFunc(cur) {
			return i
		}
	}
	return -1
}
