package tckimlik

import (
	"math/rand"
	"time"
)

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max- min) +min
	return  num
}
