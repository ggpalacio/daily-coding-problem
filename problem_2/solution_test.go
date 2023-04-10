package problem_2_test

import (
	"github.com/ggpalacio/daily-coding-problem/problem_2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int{120, 60, 40, 30, 24}, problem_2.Solution([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{2, 3, 6}, problem_2.Solution([]int{3, 2, 1}))

}
