package main

import (
	"fmt"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
)

/*****************************************
	Case Code
******************************************/

func runCase(manager *manager.GCJManager) (result string) {

	var N uint64
	fmt.Scanf("%d", &N)

	var y, maxDif, z uint64

	cache := make([]uint64, N)
	fmt.Scanf("%d", &(cache[0]))
	for i := uint64(1); i < N; i++ {
		fmt.Scanf("%d", &(cache[i]))
		if cache[i] >= cache[i-1] {
			continue
		}
		dif := cache[i-1] - cache[i]
		y += dif
		if dif > maxDif {
			maxDif = dif
		}
	}

	manager.InputUnlock()

	for i := uint64(0); i < N-1; i++ {
		dif := cache[i]
		if dif > maxDif {
			z += maxDif
		} else {
			z += dif
		}
	}

	return fmt.Sprintf("%d %d", y, z)

}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	var cases uint64
	fmt.Scanf("%d", &cases)

	output := manager.NewGCJManager(cases, runCase)
	output.WaitEnd()
}
