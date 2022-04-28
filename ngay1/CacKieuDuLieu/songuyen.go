/* Các kiểu số nguyên trong Go là:
	- uint8, uint16, uint32, uint64, int8, int16, int32, int64, int
	- uint tức là unsigned int – là kiểu số nguyên không âm
	- byte là tên gọi khác của uint8
	- rune là tên gọi khác của int32
Giới hạn:
	uint8	0 – 255
	uint16	0 – 65535
	uint32	0 – 4294967295
	uint64	0 – 18446744073709551615
	int8	-128 – 127
	int16	-32768 – 32767
	int32	-2147483648 – 2147483647
	int64	-9223372036854775808 – 9223372036854775807
*/

package main

import "fmt"

func main() {
	var num1 int16 = 20132
	var num2 int16 = 23244
	fmt.Println(num1 + num2)

	var num3 int32 = 20132
	var num4 int32 = 23244
	fmt.Println(num3 + num4)

	var num5 int = 20132
	var num6 int = 23244
	fmt.Println(num5 + num6)
}
