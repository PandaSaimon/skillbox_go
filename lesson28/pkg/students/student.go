package students_v2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s *Student) SetValues(name string, age int, grade int) {
	s.name, s.age, s.grade = name, age, grade
}

func (s Student) Print() {
	fmt.Println(s)
}

func (s Student) GetName() string {
	return s.name
}

func (s Student) GetAge() int {
	return s.age
}

func (s Student) GetGrade() int {
	return s.grade
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) SetAge(age int) {
	s.age = age
}

func (s *Student) SetGrade(grade int) {
	s.grade = grade
}

func SetValuesFromString(input string) (s Student, err error) {
	err = nil
	vals := strings.Split(input, " ")
	if len(vals) < 2 {
		err = errors.New("Invalid arguments")
		return
	}
	age, e := strconv.Atoi(vals[1])
	if e != nil {
		err = e
		return
	}
	grade, e := strconv.Atoi(vals[2])
	if e != nil {
		err = e
		return
	}
	s.SetValues(vals[0], age, grade)
	return
}
