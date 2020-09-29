package utils

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// GetIO returns a string with what the user typed in console
func GetIO() string {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// GetRandom returns a random number up to the max you specifiy
func GetRandom(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
