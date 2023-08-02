package main

import "log"

func DeleteByIndexV3[T any](index int, s []T) []T {
	if index < 0 || index >= len(s) {
		log.Fatal("Invalid index!")
	}

	// Remove the element at index
	copy(s[index:], s[index+1:])
	s = s[:len(s)-1]

	// Check if it's worth shrinking the capacity
	if len(s) <= cap(s)/2 {
		newSlice := make([]T, len(s))
		copy(newSlice, s)
		return newSlice
	}

	return s
}
