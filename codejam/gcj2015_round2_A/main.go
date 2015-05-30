package main

import (
	"flag"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
	"github.com/tiagofalcao/GoNotebook/log"
)

type position struct {
	d rune
	c bool
}

const (
	blankArrow uint8 = iota
	upArrow
	downArrow
	leftArrow
	rightArrow
	failArrow
)

var arrowRune = [...]rune{'∘', '⇧', '⇩', '⇦', '⇨', '✣'}

func printG(grid [][]uint8) {

	for i := int(0); i < len(grid); i++ {
		for j := int(0); j < len(grid[i]); j++ {
			log.Info.Printf("%c", arrowRune[grid[i][j]])
		}
		log.Info.Printf("\n")
	}
	log.Info.Printf("\n")
}

/*****************************************
	Case Code
******************************************/

func runCase(manager *manager.GCJManager) (result string) {

	var R, C uint64
	fmt.Fscanf(manager.Input, "%d %d\n", &R, &C)

	grid := make([][]uint8, R)
	for i := uint64(0); i < R; i++ {
		grid[i] = make([]uint8, C)

		for j := uint64(0); j < C; j++ {
			r, _, _ := manager.Input.ReadRune()
			switch r {
			case '.':
				grid[i][j] = blankArrow
			case '^':
				grid[i][j] = upArrow
			case 'v':
				grid[i][j] = downArrow
			case '<':
				grid[i][j] = leftArrow
			case '>':
				grid[i][j] = rightArrow
			default:
			}
		}
		manager.Input.ReadRune()
	}

	manager.InputUnlock()

	if log.InfoEnabled {
		printG(grid)
	}

	var changes uint64

	safe := make([][]bool, R)
	for i := uint64(0); i < R; i++ {
		safe[i] = make([]bool, C)
	}

	for i := uint64(0); i < R; i++ {
		for j := uint64(0); j < C; j++ {
			if grid[i][j] == blankArrow {
				continue
			} else if safe[i][j] {
				continue
			}
			if log.DebugEnabled {
				printG(grid)
			}

			switch grid[i][j] {
			case leftArrow:
				for ij := int(j) - 1; ij >= 0; ij-- {
					if grid[i][ij] != blankArrow {
						safe[i][j] = true
						break
					}
				}
			case rightArrow:
				for ij := int(j) + 1; ij < int(C); ij++ {
					if grid[i][ij] != blankArrow {
						safe[i][j] = true
						break
					}
				}
			case upArrow:
				for ii := int(i) - 1; ii >= 0; ii-- {
					if grid[ii][j] != blankArrow {
						safe[i][j] = true
						break
					}
				}
			case downArrow:
				for ii := int(i) + 1; ii < int(R); ii++ {
					if grid[ii][j] != blankArrow {
						safe[i][j] = true
						break
					}
				}
			}

			if safe[i][j] {
				continue
			}

			if grid[i][j] != leftArrow {
				for ij := int(j) - 1; ij >= 0; ij-- {
					if grid[i][ij] != blankArrow {
						grid[i][j] = leftArrow
						safe[i][j] = true
						changes++
						break
					}
				}
				if safe[i][j] {
					continue
				}
			}

			if grid[i][j] != rightArrow {
				for ij := int(j) + 1; ij < int(C); ij++ {
					if grid[i][ij] != blankArrow {
						grid[i][j] = rightArrow
						safe[i][j] = true
						changes++
						break
					}
				}
				if safe[i][j] {
					continue
				}
			}

			if grid[i][j] != upArrow {
				for ii := int(i) - 1; ii >= 0; ii-- {
					if grid[ii][j] != blankArrow {
						grid[i][j] = upArrow
						safe[i][j] = true
						changes++
						break
					}
				}
				if safe[i][j] {
					continue
				}
			}

			if grid[i][j] != downArrow {
				for ii := int(i) + 1; ii < int(R); ii++ {
					if grid[ii][j] != blankArrow {
						grid[i][j] = downArrow
						safe[i][j] = true
						changes++
						break
					}
				}
				if safe[i][j] {
					continue
				}
			}

			if log.InfoEnabled {
				grid[i][j] = failArrow
				printG(grid)
			}
			return "IMPOSSIBLE"
		}
	}

	if log.InfoEnabled {
		printG(grid)
	}
	return fmt.Sprintf("%d", changes)
}

/**********************************************************
  Google Code Jam Main
***********************************************************/
func main() {
	flag.Parse()
	manager.NewGCJManager(runCase).WaitEnd()
}
