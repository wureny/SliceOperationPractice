package main

import "log"

func DeleteByIndexV2[T any](index int, s []T) []T {
	if index < 0 || index >= len(s) {
		log.Fatal("Invalid index!")
	}

	// Remove the element at index
	copy(s[index:], s[index+1:])
	s = s[:len(s)-1]
	return s
}
