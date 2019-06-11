package main

import (
	"fmt"
	"sort"
)

func main() {

	people := []struct {
		Age  int
		Score int64
	}{
		{25, 10},
		{75, 15},
		{75, 9},
		{75, 20},
		{75, 19},
		{75, 23},
		{25, 45},
		{25, 9},
		{25, 4},
		{75, 32},
		{25, 90},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool {
		//return people[i].Name < people[j].Name
		a, b := people[i], people[j]
		if a.Age < b.Age {
			return true
		} else if a.Age > b.Age {
			return false
		}

		if a.Score < b.Score {
			return true
		}
		return false
	})
	fmt.Println("By name:", people)
}

