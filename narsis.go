package main

func Narcissistic(num int) bool {

	slice := strconv.Itoa(num)
	println(slice, len(slice))
	numLen := len(slice)
	var sum int
	var n int
	for _, s := range slice {
		//println(int(s))
		n, _ = strconv.Atoi(string(s))
		sum += int(math.Pow(float64(n), float64(numLen)))
	}
	return sum == num
}

func main() {
	println(Narcissistic(153))
}
