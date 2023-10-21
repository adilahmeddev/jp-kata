package utils

func SlicesUnorderedEql[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	elems := map[T]bool{}

	for i := 0; i < len(a); i++ {
		elems[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		if !elems[b[i]] {
			return false
		}
	}

	return true
}
