package SliceOperations

import "fmt"

func testdeleteByIndex[T any]() {
	s1 := []int{1, 2, 3}
	fmt.Println("slice: [1,2,3]")
	fmt.Println("delete slice[0]")
	fmt.Println("then the outcome should be [2,3]")
	deleteByIndex(0, s1)
	fmt.Println("actually, the outcome is ", s1)
	s2 := []string{"hello", "world"}
	fmt.Println("slice: [hello,world]")
	fmt.Println("delete slice[1]")
	fmt.Println("then the outcome should be [hello]")
	deleteByIndex(1, s2)
	fmt.Println("actually, the outcome is ", s2)
}
