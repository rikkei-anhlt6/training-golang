package main

func main() {
	var age int                     // Khai báo biến ko có giá trị ban đầu, Go tự động khởi tạo nó với zero value của kiểu biến
	age = 22                        // gán
	var tuoi int = 29               // Khai báo biến tuoi kiểu int với giá trị ban đầu bằng 29
	var sotuoi = 29                 // Khai báo biến ko có kiểu dữ liệu, go dựa vào giá trị khởi tạo của biến để nhận biết nó thuộc kiểu nào
	var width, height int = 100, 50 // Khai báo nhiều biến
	var width1, height1 = 100, 50
	var width2, height2 int
	//Khai báo nhiều biến có các kiểu dữ liệu khác nhau trong một câu lệnh
	var (
		ten      = "naveen"
		chieucao = 160
		height   int
	)
	// khai báo ngắn
	count := 10
	str, so := "abc", 10

	a, b := 20, 30 // khai bao  a và b
	b, c := 40, 50 // thoả mãn vì Cú pháp rút gọn chỉ có thể được sử dụng khi ít nhất một trong các biến ở phía bên trái của: = mới được khai báo

}
