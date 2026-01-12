package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Quiz Time 1 (1)
	// divisor := findDivisor(10)
	// fmt.Println(divisor)

	//Quiz Time 1 (2)
	// extractDigit(12245)

	//Quiz Time 2 (3)
	// starTriangle1(5)

	//Quiz Time 2 (4)
	// piramid(8)

	//Quiz Time 3 (5)
	// deretAngka(5)

	//Quiz Time 3 (6)
	// deretAngka2(9)

	//Quiz Time 4 (8)
	// isPalindrome("Kasur ini rusak")

	//Quiz Time 5 (9)
	// reverseKata("tamaT")

	//Quiz Time 5 (10)
	// checkBraces("(()))))")

	//Quiz Time 6 (11)
	numPalindrome(123241)

}

// Quiz Time 1 (1)
func findDivisor(n int) []int {
	var divisor []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisor = append(divisor, i)
		}
	}
	return divisor
}

// Quiz Time 1 (2)
func extractDigit(n int) []string {
	var temp []string
	panjang := strconv.Itoa(n)
	extract := strings.Split(panjang, "")
	for i := len(panjang) - 1; i >= 0; i-- {
		temp = append(temp, extract[i])
	}
	fmt.Println(temp)
	return temp
}

// Quiz Time 2 (3)
func starTriangle1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= n-i-1 || i >= n-j-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Quiz Time 2 (4)
func piramid(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if j <= n-i+1 || i <= n-j+1 {
				fmt.Print(j)

			}
		}
		for j := 2; j <= n; j++ {
			if j <= n-i+1 || i <= n-j+1 {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}
}

// Quiz Time 3 (5)
func deretAngka(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 != 0 {
				fmt.Print(i)
			}
			if j%2 == 0 {
				fmt.Print(n - i + 1)
			}

		}
		fmt.Println()
	}
}

// Quiz Time 3 (6)
func deretAngka2(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			if i%2 != 0 {
				if j%2 != 0 {
					fmt.Print(j)
				} else {
					fmt.Print("-")
				}
			} else if i%2 == 0 {
				if j%2 == 0 {
					fmt.Print(j)
				} else {
					fmt.Print("-")
				}
			}
		}
		fmt.Println()
	}
}

// Quiz Time 4 (8)
func isPalindrome(n string) bool {
	var kanan []string
	var kalimatKiri string
	var kalimatKanan string
	var temp []string
	kata := strings.ToLower(n)
	kataPisah := strings.Split(kata, "")
	for i := 0; i < len(kata); i++ {
		kanan = append(kanan, kataPisah[i])
	}
	temp = kanan
	kalimatKanan = strings.Join(kanan, "")
	for i, j := 0, len(kata)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	kalimatKiri = strings.Join(temp, "")
	fmt.Println(kalimatKiri == kalimatKanan)
	return kalimatKiri == kalimatKanan

}

// Quiz Time 5 (9)
func reverseKata(kata string) string {
	var temp []string
	kataPisah := strings.Split(kata, "")
	var kataKiri string
	temp = kataPisah
	for i, j := 0, len(kata)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	kataKiri = strings.Join(temp, "")
	fmt.Println(kataKiri)
	return kataKiri

}

// Quiz Time 5 (10)
func checkBraces(n string) bool {
	balance := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '(' {
			balance++
		} else if n[i] == ')' {
			balance--
		}
	}
	if balance == 0 {
		fmt.Println("true")
		return true
	} else {
		fmt.Println("false")
		return false
	}
}

// Quiz Time 6 (11)
func numPalindrome(n int) bool {
	var kanan []string
	var kalimatKiri string
	var kalimatKanan string
	var temp []string
	kata := strconv.Itoa(n)
	kataPisah := strings.Split(kata, "")
	for i := 0; i < len(kata); i++ {
		kanan = append(kanan, kataPisah[i])
	}
	temp = kanan
	kalimatKanan = strings.Join(kanan, "")
	for i, j := 0, len(kata)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	kalimatKiri = strings.Join(temp, "")
	fmt.Println(kalimatKiri == kalimatKanan)
	return kalimatKiri == kalimatKanan

}
