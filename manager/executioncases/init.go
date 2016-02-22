package executioncases

import (
	"flag"
)

// Run only one test case each time
var SeqMode bool

func init() {
	flag.BoolVar(&SeqMode, "seq", false, "Force sequential mode")
}
