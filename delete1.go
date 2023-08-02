package main

import "log"

func DeleteByIndex[T any](index int, s []T) []T {
	if index < 0 || index >= len(s) {
		log.Fatal("Invalid index!")
	}
	le := len(s)
	s = append(s[0:index], s[index+1:le]...)
	//return s[:le-1]
	return s
}
