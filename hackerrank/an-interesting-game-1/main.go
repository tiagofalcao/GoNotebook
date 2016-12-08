package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
)

import "github.com/tiagofalcao/GoNotebook/manager"

//import "os"

/*****************************************
  Code
******************************************/

func runCode(r *bufio.Reader, w io.Writer) {

	var N uint32

	fmt.Fscanf(r, "%d ", &N)
	for i := uint32(0); i < N; i++ {
		var V, v, max, inv uint32

		max = 0
		inv = 0

		fmt.Fscanf(r, "%d ", &V)
		for j := uint32(0); j < V; j++ {
			fmt.Fscanf(r, "%d ", &v)
			if v > max {
				max = v
				inv++
			}
		}

		if inv%2 != 0 {
			fmt.Fprintf(w, "BOB\n")
		} else {
			fmt.Fprintf(w, "ANDY\n")
		}
	}
}

/**********************************************************
  Simple Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewManager(runCode)
}

/**********************************************************
  HackerRank Main
***********************************************************/
/*
func main() {
	runCode(bufio.NewReader(os.Stdin), os.Stdout)
}
*/
