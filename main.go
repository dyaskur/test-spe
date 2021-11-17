package main

func evenOrOdd(arr []int) int {
	var even, odd int
	for _, v := range arr {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > odd {
		return odd
	} else {
		return even
	}
}

func main() {
	arr := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	println(evenOrOdd(arr))
}
