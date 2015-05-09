package manager

import (
	"container/heap"
	"fmt"
)

// GCJManager manage the output of each GCJ case
type GCJManager struct {
	cases      uint64  // Amount of Cases
	caseTask   GCJCase // case function
	nextCase   uint64  // The Next Case
	pq         outputQueue
	inputLock  chan bool
	caseOutput chan *googleCJCase
	caseNotify chan uint64
	endNotify  chan bool
}

// NewGCJManager start the goroutine responsible by manage the output based on case
// Receive the total of cases to send a end notify after the print of last case.
// Can receive a channel to notify after each case output, to sequential exec.
func NewGCJManager(cases uint64, caseTask GCJCase) *GCJManager {
	caseOutput := make(chan *googleCJCase, 100)
	caseNotify := make(chan uint64)
	endNotify := make(chan bool)
	inputLock := make(chan bool)

	output := &GCJManager{
		cases:      cases,
		caseTask:   caseTask,
		nextCase:   1,
		pq:         make(outputQueue, 0, 1),
		inputLock:  inputLock,
		caseOutput: caseOutput,
		caseNotify: caseNotify,
		endNotify:  endNotify,
	}

	heap.Init(&output.pq)
	go output.input()
	go output.output()
	return output
}

func (output GCJManager) end() bool {
	return output.nextCase > output.cases
}

// Print send a request to print the return of a case
func (output *GCJManager) Print(caseNum uint64, value string) {
	item := &googleCJCase{
		value:   value,
		caseNum: caseNum,
	}
	output.caseOutput <- item
}

// WaitCaseEnd wait the end of one case and return the number of this case
func (output *GCJManager) waitCaseEnd() uint64 {
	return <-output.caseNotify
}

// WaitEnd wait the end of execution of all cases
func (output *GCJManager) WaitEnd() {
	<-output.endNotify
}

// InputUnlock free input to the next execution
func (output *GCJManager) InputUnlock() {
	output.inputLock <- true
}

// InputLock free input to the next execution
func (output *GCJManager) InputLock() {
	<-output.inputLock
}

func (output GCJManager) notifyCaseEnd(caseNum uint64) {
	select {
	case output.caseNotify <- caseNum:
	default:
	}
}

func (output GCJManager) printCase(item *googleCJCase) {
	fmt.Printf("Case #%d: %s\n", item.caseNum, item.value)
}

func (output *GCJManager) input() {
	for i := uint64(1); i <= output.cases; i++ {

		go func(caseNum uint64, caseManager *GCJManager) {
			value := caseManager.caseTask(output)
			output.Print(caseNum, value)
		}(i, output)

		output.InputLock()
		/*if *seqMode {
			ret := <-caseNotify
			if ret != i {
				panic("Notified another case. This is sequential mode!")
			}
		}*/
	}
}

func (output *GCJManager) output() {
	for !output.end() {
		item := <-output.caseOutput

		if item.caseNum != output.nextCase {
			heap.Push(&output.pq, item)
			continue
		}

		output.printCase(item)
		output.nextCase++
		output.notifyCaseEnd(item.caseNum)

		for output.pq.Len() > 0 && output.nextCase == output.pq.Head().caseNum {
			item := heap.Pop(&output.pq).(*googleCJCase)
			output.printCase(item)
			output.nextCase++
			output.notifyCaseEnd(item.caseNum)
		}
	}
	output.endNotify <- true
}
