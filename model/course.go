package model

import "github.com/satori/go.uuid"

type Course struct {
	ID   string
	Name string
}

type Courses struct {
	Courses []Course
}

func (c *Courses) Add(course *Course) {
	c.Courses = append(c.Courses, *course)
}

func NewCourse() *Course {
	course := Course{
		ID: uuid.NewV4().String(),
	}
	return &course
}

func NewCourses() *Courses {
	return &Courses{}
}
