package manager

import (
	"fmt"
	"io"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager/executioncases"

type FHCManager coremanager.ExecutionManager

func FHCPrint(output io.Writer, caseNum uint64, value string) {
	fmt.Fprintf(output, "Case #%d:%s", caseNum+1, value)
}

func NewFHCManager(caseTask coremanager.ExecutionCase) *coremanager.ExecutionManager {
	return coremanager.NewExecutionManager(caseTask, FHCPrint)
}
