package main

import "fmt"
import "errors"
import "time"

//type funtion func() (str *string, err error)
//
//func fun2(f funtion) (str *string, err error) {
//	return f()
//}
//
//func add(a int, b int) (str *string, err error) {
//	var s string
//	if a+b > 3 {
//		s = "3"
//		return &s, nil
//	} else {
//		s = "error"
//		return &s, errors.New("less than 3")
//	}
//}
//
//func fun1() (*string, error){
//	var (
//		str *string
//		err error
//	)
//	_, err = fun2(func() (str *string, err error) {
//		str, err = add(1, 2)
//		if err != nil {
//			return nil, err
//		}
//		return str, err
//	})
//	fmt.Println("fun1", str,"---", err)
//	return str, err
//}
//
//func main() {
//	str, err := fun1()
//	fmt.Println("main ", str,"---", err)
//}

type Info struct {
	name string
	age  int
}

type funtion func() (str string, err error)

func fun2(f funtion) (str string, err error) {
	return f()
}

func add(a int, b int, info *Info) (str string, err error) {
	var s string
	fmt.Println("info is", *info)

	if a+b > 3 {
		s = "3"
		return s, nil
	} else {
		s = "error"
		return s, errors.New("less than 3")
	}
}

func fun1() (string, error) {
	var (
		str string
		err error
	)

	info := &Info{"chenping", 28}
	_, err = fun2(func() (string, error) {
		info.name = "chenping1"
		str, err = add(1, 2, info)
		if err != nil {
			return "", err
		}
		return str, err
	})
	fmt.Println("fun1", str, "---", err, info)
	return str, err
}

func main() {
	str, err := fun1()
	fmt.Println("main ", str, "---", err)

	start := time.Now()

	now := time.Now().Nanosecond()
	time.Sleep(2 * time.Millisecond)
	end := time.Now().Nanosecond()
	timeSpend := float64(end-now) / 1000000.0

	f := func() int64 {
		return time.Now().UnixNano() % 1e6 / 1e3
	}

	fmt.Printf("end = %v, f= %v, spend %v\n", time.Since(start), f(), timeSpend)
}
