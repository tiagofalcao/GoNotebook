package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
	"github.com/tiagofalcao/GoNotebook/instances"
	"github.com/tiagofalcao/GoNotebook/log"
	"io"
	"math"
	"sort"
)

// Based from participant CCC1120

type tree struct {
	x, y int64
}

func (t tree) AngleOf(t2 tree) float64 {
	return math.Atan2(float64(t2.y-t.y), float64(t2.x-t.x))
}

type byAsc []float64

func (a byAsc) Len() int           { return len(a) }
func (a byAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAsc) Less(i, j int) bool { return a[i] < a[j] }

/*****************************************
	Case Code
******************************************/

func calc(trees []tree, i uint64) uint64 {
	N := uint64(len(trees))
	max := uint64(0)
	log.Debug.Printf("Tree(%d) %d %d\n", i, trees[i].x, trees[i].y)
	angles := make([]float64, 0, N-1)
	for k := uint64(0); k < N; k++ {
		if i != k {
			angle := trees[i].AngleOf(trees[k])
			angles = append(angles, angle)
		}
	}
	sort.Sort(byAsc(angles))
	for k := uint64(0); k < N-1; k++ {
		angles = append(angles, angles[k]+math.Pi*2)
	}
	//log.Debug.Println(angles)
	for k := uint64(0); k < N-1; k++ {
		l := k + 1
		count := uint64(1)
		for angles[l]-angles[k] <= math.Pi {
			//log.Debug.Println(angles[k], angles[l])
			count++
			l++
		}
		if count > max {
			max = count
		}
	}
	return N - 1 - max
}

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {

	var N uint64
	fmt.Fscanf(input, "%d\n", &N)
	log.Debug.Println(N)

	trees := make([]tree, N)
	for i := uint64(0); i < N; i++ {
		fmt.Fscanf(input, "%d %d\n", &trees[i].x, &trees[i].y)
	}

	inputLock <- true
	log.Debug.Println(trees)

	var b bytes.Buffer

	if N <= 3 {
		for i := uint64(0); i < N; i++ {
			fmt.Fprintf(&b, "\n0")
		}
		io.WriteString(&b, "\n")
		return b.String()
	}

	man := instances.ResultMan(N)
	for i := uint64(0); i < N; i++ {
		go func(r *instances.ResultManager, i uint64, trees []tree) {
			man.Set(i, calc(trees, i))
		}(man, i, trees)
	}
	man.Join()
	log.Debug.Println("All trees ended")

	for i := uint64(0); i < N; i++ {
		fmt.Fprintf(&b, "\n%d", man.Results[i].(uint64))
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
