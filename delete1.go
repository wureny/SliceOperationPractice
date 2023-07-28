package SliceOperations

import "log"

func deleteByIndex[T any](index int, s []T) {
	if index < 0 && index >= len(s) {
		log.Fatal("Invalid index!")
	}
	s = append(s[0:index], s[index+1:]...)
}
