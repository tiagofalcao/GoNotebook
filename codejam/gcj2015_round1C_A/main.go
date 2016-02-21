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

	var R, C, W uint64
	fmt.Fscanf(input, "%d %d %d\n", &R, &C, &W)

	inputLock <- true

	if C%W == 0 {
		return fmt.Sprintf(" %d\n", (R*(C/W))-1+W)
	}
	return fmt.Sprintf(" %d\n", (R*(C/W))+W)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
