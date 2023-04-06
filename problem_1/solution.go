package problem_1

// Solution returns true if any two numbers of a given list of numbers added up to the given number k, false otherwise.
// This solution iterates the list of numbers looking for the number in a map and if not founds put a new value in the
// map using as key the K - the number and an empty interface as value.
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
