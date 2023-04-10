package problem_2

type Numbers []int

func Solution(input []int) (output []int) {
	output = make([]int, len(input))
	numbers := Numbers(input)

	product, hasZero := numbers.Reduce()
	for index, number := range numbers {
		if hasZero {
			if number == 0 {
				output[index] = product
			} else {
				output[index] = 0
			}
		} else {
			output[index] = product / number
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
