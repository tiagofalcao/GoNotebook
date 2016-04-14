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
	var K, C, S uint64
	fmt.Fscanf(input, "%d %d %d ", &K, &C, &S)

	inputLock <- true

	if S*C < K {
		return " IMPOSSIBLE\n"
	}

	var b bytes.Buffer

	k := uint64(0)
	for k < K {
		h := uint64(0)

		for i:=k; i<k+C && i<K; i++ {
			h *= K;
			h += i;
		}
		fmt.Fprintf(&b, " %d", h+1)
		k += C;
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
