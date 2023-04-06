package problem_1

// Solution returns true if any two numbers of a given list of numbers added up to the given number k, false otherwise.
func Solution(numbers []int, k int) bool {
	var empty interface{}
	set := make(map[int]interface{})
	for _, number := range numbers {
		if _, ok := set[number]; ok {
			return true
		}
		set[k-number] = empty
	}
	return false
}
