package utils

func MapAdd[M ~map[K]V, K comparable, V any](dst, src M) {
	for k, v := range src {
		dst[k] = v
	}
}

func MapRemove[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}

func MapClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

func MapClone[M ~map[K]V, K comparable, V any](m M) M {
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

func MapCopy[M ~map[K]V, K comparable, V any](dst, src M) {
	for k, v := range src {
		dst[k] = v
	}
}

func MapEqual[M1 ~map[K]V, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func MapEqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}
