package executioncases

import (
	"flag"
	"runtime"
)

// Run only one test case each time
var ParallelMode int

func init() {
	flag.IntVar(&ParallelMode, "n", runtime.GOMAXPROCS(-1), "How many cases in parallel")
}
