package problem_2

type Numbers []int

func Solution(input []int) (output []int) {
	output = make([]int, len(input))
	numbers := Numbers(input)

	product, hasZero := numbers.Reduce()
	for index, number := range numbers {
		if !hasZero {
			output[index] = product / number
		} else if number == 0 {
			output[index] = product
		}
	}
	return output
}

func (ref Numbers) Reduce() (product int, hasZero bool) {
	product = 1
	for _, number := range ref {
		if number == 0 {
			hasZero = true
		} else {
			product *= number
		}
	}
	return
}
