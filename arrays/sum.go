package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//func SumAll(numbersToSum ...[]int) []int {
//	lenghtOfNumbers := len(numbersToSum)
//	sums := make([]int, lenghtOfNumbers)
//
//	for i, numbers := range numbersToSum {
//		sums[i] = Sum(numbers)
//	}
//
//	return sums
//}
//
