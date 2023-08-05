package main

import (
	"solution"
)

func main() {
	var initCourses []string
	for course := range solution.Prereqs {
		initCourses = append(initCourses, course)
	}

	solution.BreadthFirst(solution.PrereqsSummary, initCourses)
}
