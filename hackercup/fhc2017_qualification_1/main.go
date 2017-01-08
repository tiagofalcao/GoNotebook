package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"

	"github.com/tiagofalcao/GoNotebook/hackercup/manager"
)

/*****************************************
	Case Code
******************************************/

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {

	var P, X, Y float64
	fmt.Fscanf(input, "%f %f %f\n", &P, &X, &Y)

	inputLock <- true

	X -= 50.0
	Y = 0.0 - (50.0 - Y)

	if math.Sqrt(math.Pow(X, 2)+math.Pow(Y, 2)) > 50.0 {
		return " white\n"
	}

	p := (math.Atan2(X, Y) * 100) / (2 * math.Pi)
	if p < 0.0 {
		p += 100.0
	}

	if p <= P {
		return " black\n"
	}

	return " white\n"
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewFHCManager(runCase).WaitEnd()
}
