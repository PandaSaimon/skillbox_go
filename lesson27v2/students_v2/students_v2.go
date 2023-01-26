package students_v2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Students struct {
	Studs map[string]*student
}

func NewStudents() Students {
	studs := &Students{Studs: make(map[string]*student)}
	return *studs
}

func (s Students) Print() {
	for _, v := range s.Studs {
		v.Print()
	}
}

func (s *Students) Add(input string) error {
	st, err := SetValuesFromString(input)
	if err != nil {
		fmt.Println(err)
	} else {
		s.Studs[st.GetName()] = &st
	}
	return err
}

type student struct {
	name  string
	age   int
	grade int
}

func (s *student) SetValues(name string, age int, grade int) {
	s.name, s.age, s.grade = name, age, grade
}

func (s student) Print() {
	fmt.Println(s)
}

func (s student) GetName() string {
	return s.name
}

func (s student) GetAge() int {
	return s.age
}

func (s student) GetGrade() int {
	return s.grade
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s *student) SetAge(age int) {
	s.age = age
}

func (s *student) SetGrade(grade int) {
	s.grade = grade
}

func SetValuesFromString(input string) (s student, err error) {
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
