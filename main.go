package main

import "fmt"

func main() {
	fmt.Println("测试delete是否成功")
	fmt.Println("V1版本：")
	DeleteByIndexTest()
	fmt.Println("V2版本：")
	DeleteByIndexTestV2()
	fmt.Println("V3版本：")
	DeleteByIndexTestV3()
}
