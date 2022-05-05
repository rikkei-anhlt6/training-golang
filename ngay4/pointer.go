package main

import "fmt"

func chaneValue(val *int) {
	*val = 100
}
func returnPointer() *int {
	i := 5
	return &i
}

func modifyArr(arr *[3]int) {
	(*arr)[0] = 90 // có thể viết tắt thành arr[0] vẫn đúng
}
func main() {
	b := 5
	var a *int = &b // gán địa chỉ của b cho a có kiểu *int
	//var c *int       // zero value là nil
	size := new(int) // tạo con trỏ kiểu int
	*size = 85       // gán giá trị con trỏ
	fmt.Println(*size)
	*a = 10 // thay đổi giá trị của b qua a
	fmt.Println(b)
	var d = 1000
	chaneValue(&d) // thay đổi giá trị qua con trỏ
	fmt.Println(d)

	e := returnPointer()
	fmt.Println(e)

	arr := [3]int{89, 90, 91}
	modifyArr(&arr)
	fmt.Println(arr)
}
