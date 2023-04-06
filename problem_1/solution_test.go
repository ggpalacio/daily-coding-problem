package problem_1_test

import (
	"github.com/ggpalacio/daily-coding-problem/problem_1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	numbers := []int{10, 15, 3, 7}

	assert.True(t, problem_1.Solution(numbers, 17))
	assert.False(t, problem_1.Solution(numbers, 15))
	assert.True(t, problem_1.Solution(numbers, 10))
}
