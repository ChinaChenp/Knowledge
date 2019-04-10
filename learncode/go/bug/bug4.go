package main

import "fmt"

type T1 struct {
	Name string
	ID   int64
	Type int64
}

type T2 struct {
	ID   int64
	Type *int64
}

func main() {
	sourceList := []T1{
		{Name: "name1", ID: 1, Type: 1},
		{Name: "name2", ID: 2, Type: 2},
		{Name: "name3", ID: 3, Type: 3},
	}
	fmt.Println(sourceList)
	retMap := map[int64]T2{}
	for _, v := range sourceList {
		fmt.Printf("v %+v, info %v\n", &v, v)
		temp := T2{
			ID:   v.ID,
			Type: &v.Type,
		}
		retMap[v.ID] = temp
	}

	fmt.Printf("retMap %+v\n", retMap)
	for _, retMapV := range retMap {
		fmt.Println(*retMapV.Type)
	}
}
