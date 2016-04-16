package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
)

/*****************************************
	Case Code
******************************************/

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {
	var N int
	fmt.Fscanf(input, "%d ", &N)

	var C [2500]int

	for i:=1; i<(2*N); i++ {
		for j:=0; j<N; j++{
			var v int
			fmt.Fscanf(input, "%d ", &v)
			C[v-1] += 1
		}
	}

	inputLock <- true

	var b bytes.Buffer
	for i:=0 ; i<len(C); i++ {
		if C[i] % 2 == 1 {
			fmt.Fprintf(&b, " %d", i+1)
			}
	}

	io.WriteString(&b, "\n")

	return b.String()
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
