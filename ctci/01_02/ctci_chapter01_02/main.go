package main

import (
	"flag"

	solution "github.com/tiagofalcao/GoNotebook/ctci/01_02"
	"github.com/tiagofalcao/GoNotebook/ctci/manager"
)

func main() {
	flag.Parse()
	manager.NewManager(solution.RunCase).WaitEnd()
}
