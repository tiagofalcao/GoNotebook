package manager

import (
	"fmt"
	"io"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"

type ExecutionManager coremanager.ExecutionManager

func Print(output io.Writer, caseNum uint64, value string) {
	fmt.Fprintf(output, "Case #%d:%s", caseNum+1, value)
}

func NewManager(caseTask coremanager.ExecutionCase) *coremanager.ExecutionManager {
	return coremanager.NewExecutionManager(caseTask, Print)
}
