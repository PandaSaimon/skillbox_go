package storage

import (
	"fmt"
	student "student/pkg/students"
)

type Students struct {
	Studs map[string]*student.Student
}

func NewStudents() Students {
	studs := &Students{Studs: make(map[string]*student.Student)}
	return *studs
}

func (s Students) Print() {
	for _, v := range s.Studs {
		v.Print()
	}
}

func (s *Students) Add(input string) error {
	st, err := student.SetValuesFromString(input)
	if err != nil {
		fmt.Println(err)
	} else {
		s.Studs[st.GetName()] = &st
	}
	return err
}
