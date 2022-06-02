package coverage

import (
	"os"
	"testing"
	"time"
)

var Alex_Sold Person = Person{firstName: "Alex", lastName: "Sold", birthDay: time.Date(2002, 3, 4, 10, 30, 0, 0, time.UTC)}
var Alla_Gulf Person = Person{firstName: "Alla", lastName: "Gulf", birthDay: time.Date(2002, 3, 4, 10, 40, 0, 0, time.UTC)}
var Alex_Ford Person = Person{firstName: "Alex", lastName: "Ford", birthDay: time.Date(2004, 6, 8, 11, 20, 0, 0, time.UTC)}

const error_sort string = "wrong sorting"

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

func TestLen(t *testing.T) {
	var people People
	people = append(people, Alex_Sold)
	res := people.Len()
	if res != 1 {
		t.Errorf("invalid input length, expected result '1'")
	}
}

func TestSwap(t *testing.T) {
	var people People
	people = append(people, Alex_Sold)
	people = append(people, Alla_Gulf)
	people.Swap(0, 1)
	if people[1].firstName != "Alex" {
		t.Errorf("elements are not swapped")
	}
}

func TestLessSameBirthDay(t *testing.T) {
	var people People
	people = append(people, Alex_Sold)
	people = append(people, Alla_Gulf)
	res := people.Less(1, 0)
	if !res {
		t.Errorf(error_sort)
	}
}

func TestLessSameFirstName(t *testing.T) {
	var people People
	people = append(people, Alex_Sold)
	people = append(people, Alex_Ford)
	res := people.Less(1, 0)
	if !res {
		t.Errorf(error_sort)
	}
}

func TestLessAllDifferent(t *testing.T) {
	var people People
	people = append(people, Alex_Ford)
	people = append(people, Alla_Gulf)
	res := people.Less(0, 1)
	if !res {
		t.Errorf(error_sort)
	}
}
