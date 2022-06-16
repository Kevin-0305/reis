package utils

// type SLICE[V any] []V

type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

//slice 切片操作
//判断slice是否值相同
func SliceEqual[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

//slice复制
func SliceCopy[E comparable](s []E) []E {
	return append([]E{}, s...)
}

//slice删除
func SliceDelete[E comparable](s []E, index int) []E {
	return append(s[:index], s[index+1:]...)
}

//slice插入
func SliceInsert[E comparable](s []E, index int, v E) []E {
	return append(append(s[:index], v), s[index:]...)
}

//slice移除值
func SliceRemoveValue[E comparable](s []E, v E) []E {
	for i, val := range s {
		if val == v {
			s = SliceDelete(s, i)
		}
	}
	return s
}

//slice替换
func SliceReplace[E comparable](s []E, index int, v E) []E {
	s[index] = v
	return s
}

//slice插入
func SliceInsertSlice[E comparable](s []E, index int, v []E) []E {
	return append(append(s[:index], v...), s[index:]...)
}

//slice排序
// func SliceSort[E comparable](s []E) []E {
// 	sor

// }

//slice筛选
func SliceFilter[E any](s []E, f func(E) bool) []E {
	var r []E
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

//快速排序
func SliceQuickSort[E Ordered](s []E) []E {
	if len(s) <= 1 {
		return s
	}
	pivot := s[0]
	left, right := []E{}, []E{}
	for _, v := range s[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(SliceQuickSort(left), pivot), SliceQuickSort(right)...)
}

// 快速排序指定函数
func SliceQuickSortFunc[E comparable](s []E, f func(E, E) bool) []E {
	if len(s) <= 1 {
		return s
	}
	pivot := s[0]
	left, right := []E{}, []E{}
	for _, v := range s[1:] {
		if f(v, pivot) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(SliceQuickSortFunc(left, f), pivot), SliceQuickSortFunc(right, f)...)
}

//快速排序指定key
func SliceQuickSortKey[E comparable, T Ordered](s []E, key func(E) T) []E {
	if len(s) <= 1 {
		return s
	}
	pivot := s[0]
	left, right := []E{}, []E{}
	for _, v := range s[1:] {
		if key(v) <= key(pivot) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(SliceQuickSortKey(left, key), pivot), SliceQuickSortKey(right, key)...)
}
