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

	var C, D, V uint64
	fmt.Fscanf(manager.Input, "%d %d %d", &C, &D, &V)

	values := make([]uint64, D)
	for i := uint64(0); i < D; i++ {
		fmt.Fscanf(manager.Input, "%d", &values[i])
	}

	manager.InputUnlock()

	var count uint64
	var sum uint64
	var i uint64

	if values[0] != 1 {
		count++
		//fmt.Printf("New %d [%d]\n", 1, sum)
	} else {
		i = uint64(1)
	}
	sum += C

	for sum < V {
		if i >= uint64(len(values)) {
			//values = append(values, sum+1)
			count++
			//fmt.Printf("New %d [%d]\n", sum+1, sum)
			sum += (sum + 1) * C
			continue
		} else {
			for sum < (values[i]-1) && sum < V {
				count++
				//fmt.Printf("New %d [%d]\n", sum+1, sum)
				sum += (sum + 1) * C
			}
		}
		//fmt.Printf("Used %d [%d]\n", values[i], sum)
		sum += values[i] * C
		i++
	}

	//fmt.Printf("End [%d]\n", sum)
	return fmt.Sprintf("%d", count)

}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
