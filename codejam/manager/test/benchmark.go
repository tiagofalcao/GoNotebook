package test

import (
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
	"testing"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"
import coretest "github.com/tiagofalcao/GoNotebook/manager/executioncases/test"

func Benchmark(b *testing.B, input string, caseTask coremanager.ExecutionCase) {
	coretest.Benchmark(b, input, caseTask, manager.GCJPrint)
}
