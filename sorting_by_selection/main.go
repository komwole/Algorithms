package main
import "fmt"
func main() {
	var arr = []int{1, 3, 2, 2, 4, -8, -4, 10}
	length := len(arr)
	fmt.Println(selection_sort(arr, length))
}

func find_smallest(list []int) int{
	smallest := list[0]
	smallest_index := 0
	for i := 1; i < len(list); i++ {
		if list[i] < smallest {
			smallest = list[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selection_sort(list []int, length int) []int {
	var newlist []int
	for i := 0; i < length; i++ {
		smallest := find_smallest(list)
		newlist = append(newlist, list[smallest])
		list = append(list[:smallest], list[smallest+1:]...)
	}
	return newlist
}
