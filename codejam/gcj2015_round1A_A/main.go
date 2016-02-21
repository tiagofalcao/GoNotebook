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

	var N uint64
	fmt.Fscanf(input, "%d\n", &N)

	var y, maxDif, z uint64

	cache := make([]uint64, N)
	fmt.Fscanf(input, "%d ", &(cache[0]))
	for i := uint64(1); i < N; i++ {
		fmt.Fscanf(input, "%d ", &(cache[i]))
		if cache[i] >= cache[i-1] {
			continue
		}
		dif := cache[i-1] - cache[i]
		y += dif
		if dif > maxDif {
			maxDif = dif
		}
	}

	inputLock <- true

	for i := uint64(0); i < N-1; i++ {
		dif := cache[i]
		if dif > maxDif {
			z += maxDif
		} else {
			z += dif
		}
	}

	return fmt.Sprintf(" %d %d\n", y, z)

}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
