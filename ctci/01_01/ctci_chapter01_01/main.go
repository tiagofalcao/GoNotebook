package main

import (
	"flag"

	solution "github.com/tiagofalcao/GoNotebook/ctci/01_01"
	"github.com/tiagofalcao/GoNotebook/ctci/manager"
)

func main() {
	flag.Parse()
	manager.NewManager(solution.RunCase).WaitEnd()
}
