package testbench

import (
	"log"
	"testing"
)

func DeleteByIndex[T any](index int, s []T) []T {
	if index < 0 || index >= len(s) {
		log.Fatal("Invalid index!")
	}
	le := len(s)
	s = append(s[0:index], s[index+1:le]...)
	//return s[:le-1]
	return s
}
func DeleteByIndexV2[T any](index int, s []T) []T {
	if index < 0 || index >= len(s) {
		log.Fatal("Invalid index!")
	}

	// Remove the element at index
	copy(s[index:], s[index+1:])
	s = s[:len(s)-1]
	return s
}
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
func BenchmarkDeleteByIndex(b *testing.B) {
	s1 := make([]int, 0, 7)
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7)
	for n := 0; n < b.N; n++ {
		s2 := make([]int, len(s1), cap(s1))
		s2 = DeleteByIndex(0, s2)
	}
}
func BenchmarkDeleteByIndexV2(b *testing.B) {
	s1 := make([]int, 0, 7)
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7)
	for n := 0; n < b.N; n++ {
		s2 := make([]int, len(s1), cap(s1))
		s2 = DeleteByIndexV2(0, s2)
	}
}
func BenchmarkDeleteByIndexV3(b *testing.B) {
	s1 := make([]int, 0, 7)
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7)
	for n := 0; n < b.N; n++ {
		s2 := make([]int, len(s1), cap(s1))
		s2 = DeleteByIndexV3(0, s2)
	}
}
