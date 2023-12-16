package main

import "fmt"

func gcd(x, y int32) int32 {
	var c int32
	for y != 0 {
		c = x % y
		x, y = y, c
	}
	return x
}

func gcdMultiple(numbers []int32) int32 {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for _, num := range numbers[1:] {
		result = gcd(result, num)
	}
	return result
}

func lcm(x, y int32) int32 {
	var c int32
	c = x * y / gcd(x, y)
	return c
}

func lcmMultiple(numbers []int32) int32 {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for _, num := range numbers[1:] {
		result = lcm(result, num)
	}
	return result
}

func getTotalX(a []int32, b []int32) int32 {
	lcmA := lcmMultiple(a)
	fmt.Printf("lcmA : %d\n", lcmA)
	gcdB := gcdMultiple(b)
	fmt.Printf("gcdB : %d\n", gcdB)
	// นับจำนวนระหว่าง gcdB และ lcmA
	count := int32(0)
	for multiple := lcmA; multiple <= gcdB; multiple += lcmA {
		if gcdB%multiple == 0 {
			count++
		}
	}
	return count
}

func main() {
	var a []int32
	a = []int32{2, 4}
	var b []int32
	b = []int32{16, 32, 96}
	fmt.Printf("Count : %d\n", getTotalX(a, b))
}
