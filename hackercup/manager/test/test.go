package test

import (
	"testing"

	"github.com/tiagofalcao/GoNotebook/hackercup/manager"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"
import coretest "github.com/tiagofalcao/GoNotebook/manager/executioncases/test"

func Test(t *testing.T, input string, caseTask coremanager.ExecutionCase, output string) {
	coretest.Test(t, input, caseTask, manager.FHCPrint, output, nil)
}
