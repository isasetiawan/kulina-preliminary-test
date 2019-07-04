package algoritma

import (
	"fmt"
	"strings"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimesNumber(n int) []int {
	var temp []int
	i := 0
	for len(temp) <= n {
		if IsPrime(i) {
			temp = append(temp, i)
		}
		i += 1
	}
	return temp
}

func Fibonaci(n int) []int {
	temp := make([]int, n+1, n+2)
	if n < 2 {
		temp = temp[0:2]
	}
	temp[0] = 0
	temp[1] = 1
	for i := 2; i <= n; i++ {
		temp[i] = temp[i-1] + temp[i-2]
	}
	return temp

}

func Haha(numbers string) {
	n := len(numbers)
	for i, number := range numbers {
		fmt.Printf(" %c%s \n", number, strings.Repeat("0", n-i-1))
	}
}
