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
	fmt.Fscanf(input, "%d ", &N)

	inputLock <- true

	if N == 0 {
		return " INSOMNIA\n"
	}

	D := [10]bool {false,false,false,false,false,false,false,false,false,false}
	Dn := uint64(10)

	M := N
	C := uint64(1)
	for {
		K := M
		for K > 0 {
			k := K % 10
			if !D[k] {
				D[k] = true
				Dn -= 1
				if Dn == 0 {
					return fmt.Sprintf(" %d\n", M)
				}
			}
			K /= 10
		}
		M += N
		C += 1
	}
	return " INSOMNIA\n"
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
