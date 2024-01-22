package problem_399_test

import (
	"github.com/ggpalacio/daily-coding-problem/problem_399"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, [][]uint{{3, 5}, {8, 0}, {8}}, problem_399.Solution([]uint{3, 5, 8, 0, 8}))
	assert.Equal(t, [][]uint{{8}, {3, 5, 0}, {8}}, problem_399.Solution([]uint{8, 3, 5, 0, 8}))
	assert.Equal(t, [][]uint{{8}, {3, 5}, {8, 0}}, problem_399.Solution([]uint{8, 3, 5, 8, 0}))
	assert.Equal(t, [][]uint{{3, 5}, {1, 7}, {8, 0}}, problem_399.Solution([]uint{3, 5, 1, 7, 8, 0}))
	assert.Equal(t, [][]uint{{6}, {1, 1, 1, 1, 1, 1}, {3, 3}}, problem_399.Solution([]uint{6, 1, 1, 1, 1, 1, 1, 3, 3}))
	assert.Equal(t, [][]uint{{1, 1, 1, 1, 1, 1}, {3, 3}, {1, 5}}, problem_399.Solution([]uint{1, 1, 1, 1, 1, 1, 3, 3, 1, 5}))
	assert.Equal(t, [][]uint{{1, 1, 1, 1, 1, 1}, {3, 3}, {6}}, problem_399.Solution([]uint{1, 1, 1, 1, 1, 1, 3, 3, 6}))
	assert.Nil(t, problem_399.Solution([]uint{8, 3, 5, 1, 8}))
}
