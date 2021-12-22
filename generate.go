package tckimlik

import (
	"math/rand"
	"strconv"
)

func Generate () string  {
	var (
		digits [11]int
		tcIdentity string
		evenNumbersTotal int
		oddNumbersTotal int
	)

	// generate the first number (1 - 10)
	digits[0] = RandomNumber(1,10)

	// generate other numbers (2-3-4-5-6-7-8)
	for i:= 1; i<9; i++ {
		digits[i] = rand.Intn(10-0) + 0
	}

	// Calculate the sum of even and odd numbers
	for i:= 1; i<9; i++ {
		if i%2==0 {
			evenNumbersTotal += digits[i]
		} else {
			oddNumbersTotal += digits[i]
		}
	}

	// the tenth digit is calculated
	digits[9] = (evenNumbersTotal*7 - oddNumbersTotal) % 10

	// the eleventh digit is calculated
	digits[10] = (evenNumbersTotal + oddNumbersTotal + digits[9]) % 10

	for _, digit := range digits {
		tcIdentity += strconv.Itoa(digit)
	}

	return tcIdentity
}