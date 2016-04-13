package manager

import (
	"fmt"
	"io"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"

type GCJManager coremanager.ExecutionManager

func GCJPrint(output io.Writer, caseNum uint64, value string) {
	fmt.Fprintf(output, "Case #%d:%s", caseNum + 1, value)
}

func NewGCJManager(caseTask coremanager.ExecutionCase) *coremanager.ExecutionManager {
	return coremanager.NewExecutionManager(caseTask, GCJPrint)
}
