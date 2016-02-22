package diff

import (
	"fmt"
	"io"
	"os"
)

func PrintLinesCompare(w io.Writer, output, expected []string) {
	l := len(output)
	if len(expected) < l {
		l = len(expected)
	}

	var i int = 0;

  for ; i<l; i++ {
		if output[i] == expected[i] {
			fmt.Fprintf(w, "%-35s   %-35s\n", output[i], expected[i])
		} else {
			fmt.Fprintf(w, "%-35s | %-35s\n", output[i], expected[i])
		}
  }

	for ; i<len(output); i++ {
			fmt.Fprintf(w, "%-35s | \n", output[i])
	}

	for ; i<len(expected); i++ {
			fmt.Fprintf(w, "%-35s | %-35s\n", "", expected[i])
	}
}

func CompareDefault(output, expected []string, v interface{}) bool {
	if len(output) != len(expected) {
		PrintLinesCompare(os.Stderr, output, expected)
		return false
	}

  for i:=0; i<len(output); i++ {
		if output[i] != expected[i] {
			PrintLinesCompare(os.Stderr, output, expected)
			return false
		}
  }

	return true
}
