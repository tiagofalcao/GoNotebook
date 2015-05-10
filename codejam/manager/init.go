package manager

import (
	"flag"
)

// OptOutput target the output to a file
var OptOutput string

// OptInput target the input from a file
var OptInput string

var seqMode bool

func init() {
	flag.StringVar(&OptOutput, "output", "", "Output File")
	flag.StringVar(&OptOutput, "o", "", "Output File")
	flag.StringVar(&OptOutput, "input", "", "Output File")
	flag.StringVar(&OptOutput, "i", "", "Output File")
	flag.BoolVar(&seqMode, "seq", false, "Force sequential mode")
}
