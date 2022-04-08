package slices

import "reflect"

func Contains(slice interface{}, item interface{}) int {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return -1
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return i
		}
	}
	return -1
}
