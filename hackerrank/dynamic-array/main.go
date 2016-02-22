package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/manager"
	"io"
)

/*****************************************
	Case Code
******************************************/

func runCase(r *bufio.Reader, w io.Writer) {

	var N, Q, c, x, y, lastans uint32

	lastans = 0

	fmt.Fscanf(r, "%d %d\n", &N, &Q)

	seq := make([][]uint32, N)

	for i:=uint32(0); i<N; i++ {
		seq[i] = make([]uint32, 0, 10)
	}

	for i:=uint32(0); i<Q; i++ {
		fmt.Fscanf(r, "%d %d %d\n", &c, &x, &y)
		s := (x ^ lastans) % N

		if c == 1 { // Insert y at the end of the ((x⊕lastans) mod N)th sequence.
			seq[s] = append(seq[s], y)
		} else if c == 2 { // Print the value of the (y mod size)th element of the ((x⊕lastans) mod N)th sequence.
			lastans = seq[s][y % uint32(len(seq[s]))]
			fmt.Fprintf(w, "%d\n", lastans)
		} else {
			fmt.Fprintf(w, "Unknown command %d\n", c)
		}
	}
}

/**********************************************************
  Simple Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewManager(runCase)
}
