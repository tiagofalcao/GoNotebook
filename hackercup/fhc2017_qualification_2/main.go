package main

import (
	"bufio"
	"flag"
	"fmt"
	"sort"

	"github.com/tiagofalcao/GoNotebook/hackercup/manager"
)

type IntSort []int

func (h IntSort) Len() int           { return len(h) }
func (h IntSort) Less(i, j int) bool { return h[i] > h[j] }
func (h IntSort) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

/*****************************************
	Case Code
******************************************/

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {

	var N uint64
	fmt.Fscanf(input, "%d\n", &N)

	h := make([]int, N)

	for k := uint64(0); k < N; k++ {
		fmt.Fscanf(input, "%d\n", &h[k])
	}

	inputLock <- true

	sort.Sort(IntSort(h))

	T := uint64(0)
	for len(h) > 0 {
		p := 50 / h[0]
		if 50%h[0] > 0 {
			p++
		}
		p--
		h = h[1:]
		if p > len(h) {
			break
		}
		for i := 0; i < p; i++ {
			h = h[:len(h)-1]
		}
		T++
	}

	return fmt.Sprintf(" %d\n", T)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewFHCManager(runCase).WaitEnd()
}
