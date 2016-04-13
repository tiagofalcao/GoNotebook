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
	C := uint64(0)

  last, _, _ := input.ReadRune()
	for {
		r, _, _ := input.ReadRune()
		if r == '\n' {
			break
		} else if r != last {
			last = r
			C += 1
		}
	}

	if last != '+' {
		C += 1
	}

	inputLock <- true

	return fmt.Sprintf(" %d\n", C)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
