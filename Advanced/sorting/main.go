package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

//===================================================

// This was auto convert from type ByAge []Person	type ByAge []Person
// type (
// 	ByAge  []Person
// 	ByName []Person
// )
//
// func (a ByAge) Len() int {
// 	return len(a)
// }
//
// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }
//
// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
//
// func (a ByName) Len() int {
// 	return len(a)
// }
//
// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }
//
// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

func main() {
	people := []Person{
		{"Alice", 32},
		{"Catherine", 31},
		{"Bob", 33},
	}

	fmt.Println("Unsorted by age: ", people)

	ageAscending := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	nameDescending := func(p1, p2 *Person) bool {
		return p1.Name > p2.Name
	}
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	By(ageAscending).Sort(people)
	fmt.Println("Sorted by age ascending: ", people)

	By(nameDescending).Sort(people)
	fmt.Println("Sorted by name desceding: ", people)

	By(lenName).Sort(people)
	fmt.Println("Sorted by length of name: ", people)

	//============ Sort of Slice
	stringSlice := []string{"banana", "apple", "mango", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1] // The index here choose the last character because - 1
	})
	fmt.Println("String Slice sorted by last character: ", stringSlice)
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][0] < stringSlice[j][0]
	})
	fmt.Println("String Slice sorted by first character: ", stringSlice)

	//===================================================
	// fmt.Println("Unsorted by age: ", people)
	// sort.Sort(ByAge(people))
	// fmt.Println("Sorted by age: ", people)
	// sort.Sort(ByName(people))
	// fmt.Println("Sorted by name: ", people)

	// numbers := []int{5, 90, 12, 4, 1, 2, 45}
	// sort.Ints(numbers)
	// fmt.Println(numbers)
	//
	// stringSlice := []string{"ball", "apple", "car"}
	// sort.Strings(stringSlice)
	// fmt.Println(stringSlice)
}
