package main

func main() {
	println(max([]int{1, 9, 2, 10, 0, 3, 11, 32}))
}

func max(array []int) int {
	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		} else {
			return array[1]
		}
	}

	max := max(array[1:])
	if array[0] > max {
		return array[0]
	} else {
		return max
	}
}
