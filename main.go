package main
import "fmt"
func main() {
	var arr = []int{0, 1, 2, 5, 8, 9, 12, 18, 24, 49}
	x1 := 2
	x2 := 50
	fmt.Println(binary_search(arr, x1))
	fmt.Println(binary_search(arr, x2))
}
func binary_search(list []int, item int) int{
	low := 0 
	high := len(list) - 1
	
	for low <= high {
		mid := low + high 
		guess := list[mid]
		if guess == item {
			return mid 
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1 
}

