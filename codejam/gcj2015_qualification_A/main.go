package main

import (
	"flag"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
)

/*****************************************
	Case Code
******************************************/

func runCase(manager *manager.GCJManager) (result string) {
	var S uint64
	fmt.Fscanf(manager.Input, "%d ", &S)

	var audienceStr string
	fmt.Fscanf(manager.Input, "%s\n", &audienceStr)

	manager.InputUnlock()

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
	return fmt.Sprintf("%d", friends)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
