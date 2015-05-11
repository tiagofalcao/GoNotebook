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

	var R, C, W uint64
	fmt.Fscanf(manager.Input, "%d %d %d", &R, &C, &W)

	manager.InputUnlock()

	if C%W == 0 {
		return fmt.Sprintf("%d", (R*(C/W))-1+W)
	}
	return fmt.Sprintf("%d", (R*(C/W))+W)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
