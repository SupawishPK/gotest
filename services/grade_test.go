package services

import "testing"

func TestCheckGrade(t *testing.T) {

	type testcase struct {
		name  string
		score int
		grade string
	}

	testCases := []testcase{
		{name: "A", score: 80, grade: "A"},
		{name: "B", score: 70, grade: "B"},
		{name: "C", score: 60, grade: "C"},
		{name: "D", score: 50, grade: "D"},
		{name: "F", score: 10, grade: "F"},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			grade := checkGrade(c.score)
			expected := c.grade

			if grade != expected {
				t.Errorf("get %v expected %v", grade, expected)
			}
		})
	}
}
