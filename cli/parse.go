package cli

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func ParseArgs(args []string) (numRows, numCols, init int, err error) {
	if len(args) > 3 {
		return 0, 0, 0, fmt.Errorf("usage: %s [numRows [numCols [init]]]", os.Args[0])
	}

	numericArgs := []int{}
	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil || (i <= 1 && num < 1) {
			return 0, 0, 0, fmt.Errorf("invalid input argument (%s) at the %d-th position (0-indexed)", arg, i)
		}
		numericArgs = append(numericArgs, num)
	}

	numRows, numCols, init = ParseNumericArgs(numericArgs)
	err = nil
	return
}

func AutoNumSeeds(numRows, numCols int) int {
	return 2 * int(math.Ceil(math.Sqrt(float64(numRows*numCols))))
}

func ParseNumericArgs(args []int) (numRows, numCols, init int) {
	if len(args) == 0 {
		return 10, 10, 20
	} else if len(args) == 1 {
		return args[0], args[0], AutoNumSeeds(args[0], args[0])
	} else if len(args) == 2 {
		return args[0], args[1], AutoNumSeeds(args[0], args[1])
	} else {
		return args[0], args[1], args[2]
	}
}
