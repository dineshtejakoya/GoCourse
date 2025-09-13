package advanced

import (
	"fmt"
	"sort"
)

/*
sort.Ints ->takes slice of integers
sort.strings ->takes slice of strings
*/

// p1,p2 are pointers to function which returs a boolean value
type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

// implement len,swap and less on personsorter
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

type Person struct {
	Name string
	Age  int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	numbers := []int{242, 342, 24, 35, 3}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	stringsSlice := []string{"John", "Anthony", "Basha", "Manik Basha", "Muthu", "Lingaa"}
	sort.Strings(stringsSlice)
	fmt.Println("Sorted Strings:", stringsSlice)

	//Sorting By Functions - involves custom sorting criteria, by definining functions that determine order of elements
	//in go achieved with sort.Interface

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}

	// fmt.Println("UnSorted By Age", people)
	// sort.Sort(ByAge(people))
	// fmt.Println("Sorted By Age", people)
	// sort.Sort(ByName(people))
	// fmt.Println("Sorted By Name", people)

	fmt.Println("UnSorted By Age", people)
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted By Age", people)
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(people)
	fmt.Println("Sorted By Name", people)

	//========SORT BY DESCENDING==============
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}

	By(ageDesc).Sort(people)
	fmt.Println("Sorted By Age (Descending):", people)

	//===========SORT BY LEN(NAME)
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}
	By(lenName).Sort(people)
	fmt.Println("Sorted by Len Name", people)

	//=========SORT.SLICE==============
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last character:", stringSlice)

}
