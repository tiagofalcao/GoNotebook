package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"math/big"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
	"github.com/tiagofalcao/GoNotebook/math/primes"
)

/*****************************************
	Case Code
******************************************/

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {
	var N, J uint64
	fmt.Fscanf(input, "%d %d", &N, &J)

	inputLock <- true
	V := big.NewInt(1)
	V = V.Lsh(V, uint(N-1))
  max := big.NewInt(1)
	max = max.Lsh(V, 1)
	V = V.Add(V, big.NewInt(1))

	var b bytes.Buffer
	io.WriteString(&b, "\n")
	for J > 0 {
		D := make([]*big.Int, 9)
		d := uint64(len(D))
		S := V.Text(2)

		for j := 0; j < len(D); j++ {
			z := big.NewInt(0)
			z, _ = z.SetString(S, j+2)

			D[j] = primes.LowerBigDivFast(z, big.NewInt(10000))
			if D[j].Cmp(big.NewInt(0)) > 0 {
				d -= 1
			}
		}
		if d == 0 {
			io.WriteString(&b, S)
			for j := 0; j < len(D); j++ {
				fmt.Fprintf(&b, " %s", D[j].String())
			}
			io.WriteString(&b, "\n")
			J -= 1
		}

		V = V.Add(V, big.NewInt(2))
		if V.Cmp(max) >= 0 {
			break;
		}
	}

	return b.String()
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
