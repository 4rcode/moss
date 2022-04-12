package internal

func Set[T comparable](value *T, zero, fallback T) {
	if value != nil && *value == zero {
		*value = fallback
	}
}
