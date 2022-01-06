package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testcase struct {
		name     string
		score    int
		expected string
	}

	testCases := []testcase{
		{name: "A", score: 80, expected: "A"},
		{name: "B", score: 70, expected: "B"},
		{name: "C", score: 60, expected: "C"},
		{name: "D", score: 50, expected: "D"},
		{name: "F", score: 10, expected: "F"},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			grade := checkGrade(c.score)
			expected := c.expected

			assert.Equal(t, expected, grade)

		})
	}
}
