package main

import "fmt"

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	uniqueA := unique(a)
	uniqueB := unique(b)
	intersection := findIntersection(uniqueA, uniqueB)
	union := findUnion(uniqueA, uniqueB)

	fmt.Println(uniqueA, uniqueB)
	fmt.Println(intersection)
	fmt.Println(union)
}

func unique(slice []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func findIntersection(a, b []int) []int {
	inB := make(map[int]bool)
	for _, v := range b {
		inB[v] = true
	}

	var result []int
	for _, v := range a {
		if inB[v] {
			result = append(result, v)
		}
	}
	return result
}

func findUnion(a, b []int) []int {
	combined := append(a, b...)
	return unique(combined)
}
