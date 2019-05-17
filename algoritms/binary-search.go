package main

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index, ok := find(list, 3)

	println(index, ok)
}

func find(list []int, value int) (int, bool) {
	var mid int
	var guess int

	low := 0
	high := len(list) - 1

	for ; low <= high ; {
		mid = (high + low) / 2
		guess = list[mid]

		if guess == value {
			return mid, true
		} else if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, false
}