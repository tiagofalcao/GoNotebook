package manager

import (
	"flag"
)

// OptOutput target the output to a file
var OptOutput string

// OptInput target the input from a file
var OptInput string

// Force flush output
var FlushMode bool

func init() {
	flag.StringVar(&OptOutput, "output", "", "Output File")
	flag.StringVar(&OptOutput, "o", "", "Output File")
	flag.StringVar(&OptInput, "input", "", "Output File")
	flag.StringVar(&OptInput, "i", "", "Output File")

	flag.BoolVar(&FlushMode, "flush", true, "Force flush all output")
}
