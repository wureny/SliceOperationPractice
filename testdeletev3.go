package main

import "fmt"

func DeleteByIndexTestV3() {
	s1 := make([]int, 0, 3)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(len(s1), cap(s1))
	fmt.Println("slice: [1,2,3]")
	fmt.Println("delete slice[0]")
	fmt.Println("then the outcome should be [2,3]")
	s1 = DeleteByIndexV3(0, s1)
	fmt.Println("actually, the outcome is ", s1)
	fmt.Println(len(s1), cap(s1))
	s2 := []string{"hello", "world"}
	fmt.Println("slice: [hello,world]")
	fmt.Println("delete slice[1]")
	fmt.Println("then the outcome should be [hello]")
	s2 = DeleteByIndexV3(1, s2)
	fmt.Println("actually, the outcome is ", s2)
}
