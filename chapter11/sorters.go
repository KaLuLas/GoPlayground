package chapter11

import (
	"fmt"
	"sort"
)

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	days []*day
}

func (array *dayArray) Len() int {
	return len(array.days)
}

func (array *dayArray) Less(i, j int) bool {
	dayA := array.days[i]
	dayB := array.days[j]
	return dayA.num < dayB.num
}

func (array *dayArray) Swap(i, j int) {
	array.days[i], array.days[j] = array.days[j], array.days[i]
}

type person struct {
	firstName string
	lastName  string
}

func (p person) String() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

type personSlice []*person

func (p personSlice) Len() int {
	return len(p)
}

func (p personSlice) Less(i, j int) bool {
	if p[i].lastName == p[j].lastName {
		return p[i].firstName < p[j].firstName
	}

	return p[i].lastName < p[j].lastName
}

func (p personSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TrySortDays() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	array := &dayArray{data}
	sort.Sort(array)

	fmt.Println()
	for _, day := range array.days {
		fmt.Printf("%s ", day.shortName)
	}
	fmt.Println()
}

func TrySortPersons() {
	person1 := person{"Tom", "Hades"}
	person2 := person{"Sheldon", "Cooper"}
	person3 := person{"Harry", "Potter"}
	person4 := person{"Tom", "Potter"}
	person5 := person{"Sheldon", "Potter"}
	persons := personSlice{&person1, &person2, &person3, &person4, &person5}
	//persons := personSlice([]person{person1, person2, person3}) // same power ↑ shocking
	sort.Sort(persons)
	for _, person := range persons {
		fmt.Println(person)
	}
}
