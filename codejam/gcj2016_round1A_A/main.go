package main

import (
	"bufio"
	"flag"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
)

/*****************************************
	Case Code
******************************************/

func runCase(input *bufio.Reader, inputLock chan bool) (result string) {

	var s []rune

	r, _, _ := input.ReadRune()
	s = append(s, r)

	for {
		r, _, _ = input.ReadRune()
		if r == '\n' {
			break
		} else if r >= s[0] {
			s = append([]rune{r}, s...)
		} else {
			s = append(s, r)
		}
	}

	inputLock <- true

	s = append([]rune{' '}, s...)
	s = append(s, '\n')
	return string(s)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
