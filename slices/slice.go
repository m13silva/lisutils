package slices

import lisette "github.com/ivov/lisette/prelude"

func Last[T any](slice []T) (T, bool) {
	len_ := len(slice)
	if len_ == 0 {
		return *new(T), false
	}
	v_1 := lisette.SliceGet(slice, len_-1)
	if v_1.Tag == lisette.OptionSome {
		return v_1.SomeVal, true
	}
	return *new(T), false
}

func NotIn[T comparable](value T, options []T) bool {
	for _, opt := range options {
		if value == opt {
			return false
		}
	}
	return true
}
