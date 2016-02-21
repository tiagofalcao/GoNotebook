package test

import (
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
	"testing"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"
import coretest "github.com/tiagofalcao/GoNotebook/manager/executioncases/test"

func Test(t *testing.T, input string, caseTask coremanager.ExecutionCase, output string) {
	coretest.Test(t, input, caseTask, manager.GCJPrint, output)
}
