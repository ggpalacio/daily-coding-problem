package problem_2

type Numbers []int

func Solution(input []int) (output []int) {
	output = make([]int, len(input))
	numbers := Numbers(input)

	productAllNumbers := numbers.ProductAll()
	for index, number := range numbers {
		output[index] = productAllNumbers / number
	}
	return output
}

func (ref Numbers) ProductAll() int {
	productAll := 1
	for _, number := range ref {
		productAll *= number
	}
	return productAll
}
