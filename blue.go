package main

func blueOcean(arr1 []int, arr2 []int) []int {
	for _, v := range arr2 {
		for i, v2 := range arr1 {
			if v == v2 {
				arr1 = append(arr1[:i], arr1[i+1:]...)
			}
		}
	}
	return arr1
}
func main() {
	//printLn(blueOcean([1, 2, 3, 4, 6, 10], [1]))
	//printLn(blueOcean([1, 5, 5, 5, 5, 3], [5]))
}
