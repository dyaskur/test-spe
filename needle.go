package main

import (
	"fmt"
)

func findNeedle(haystack []string, needle string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	needle := "blue"
	fmt.Println(findNeedle(haystack, needle))
}
