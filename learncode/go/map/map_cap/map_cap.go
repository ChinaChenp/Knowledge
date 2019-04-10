package map_cap

func Func1() {
	m := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		m[i] = i * 10
	}
}

func Func2() {
	m := make(map[int]int, 1000000)
	for i := 0; i < 1000000; i++ {
		m[i] = i * 10
	}
}

//func main() {
//	m := make(map[int]int, 10)
//	for i := 0; i < 15; i++ {
//		m[i] = i * 10
//	}
//
//	cout := 0
//	for k, v := range m {
//		cout++
//		fmt.Println(k, v)
//	}
//	fmt.Println(cout)
//}
