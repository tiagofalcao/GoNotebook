package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
)

/*****************************************
	Case Code
******************************************/

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {
	var S uint64
	fmt.Fscanf(input, "%d ", &S)

	var audienceStr string
	fmt.Fscanf(input, "%s\n", &audienceStr)

	inputLock <- true

	runes := []rune(audienceStr)

	var friends, stoodUp uint64
	for i, d := range runes {
		I := uint64(i)
		audience := uint64(d - '0')

		if I > stoodUp {
			friends += I - stoodUp
			stoodUp = I
		}

		stoodUp += audience
	}
	return fmt.Sprintf(" %d\n", friends)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
