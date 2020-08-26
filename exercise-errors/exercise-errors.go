package main

import (
	"fmt"
	"math"
	"strconv"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %s",
		strconv.FormatFloat(float64(e), 'f', -1, 64))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, &err
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2.9))
	fmt.Println(Sqrt(-2.9))
}
