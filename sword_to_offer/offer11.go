package 剑指offer

func minArray(numbers []int) int {

	length := len(numbers)
	if length == 0 {
		return -1
	}

	left, right := 0, length-1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return numbers[left]
}
