package main

import "fmt"

type SumCalculator interface {
	DisplaySum()
}
type SubCalculator interface {
	CalSub() int
}
type Calculator interface {
	SumCalculator
	SubCalculator
}
type Init struct {
	a int
	b int
}

func (i Init) DisplaySum() {
	fmt.Printf("Tổng 2 số là: %d", i.a+i.b)
}

func (i Init) CalSub() int {
	return i.a - i.b
}
func main() {
	i := Init{
		a: 10,
		b: 5,
	}
	var Calculator = i
	Calculator.DisplaySum()
	fmt.Println()
	fmt.Printf("Hiệu 2 số là: %d", Calculator.CalSub())

}
