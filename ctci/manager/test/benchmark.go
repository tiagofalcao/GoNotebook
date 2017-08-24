package test

import (
	"testing"

	"github.com/tiagofalcao/GoNotebook/ctci/manager"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"
import coretest "github.com/tiagofalcao/GoNotebook/manager/executioncases/test"

func Benchmark(b *testing.B, input string, caseTask coremanager.ExecutionCase) {
	coretest.Benchmark(b, input, caseTask, manager.Print)
}
