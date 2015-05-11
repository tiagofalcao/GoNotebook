package manager

import (
	"container/heap"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/log"
	"io"
	"os"
)

// GCJManager manage the io of each GCJ case
type GCJManager struct {
	cases      uint64  // Amount of Cases
	caseTask   GCJCase // case function
	nextCase   uint64  // The Next Case
	pq         outputQueue
	inputLock  chan bool
	caseOutput chan *googleCJCase
	caseNotify chan uint64
	endNotify  chan bool
	Input      io.Reader
	Output     io.Writer
}

// NewGCJManagerIO start the goroutine responsible by manage the io based on case
func NewGCJManagerIO(caseTask GCJCase, input io.Reader, output io.Writer) *GCJManager {
	log.Init(log.DefaultLevel)

	caseOutput := make(chan *googleCJCase, 100)
	caseNotify := make(chan uint64)
	endNotify := make(chan bool)
	inputLock := make(chan bool)

	manager := &GCJManager{
		caseTask:   caseTask,
		nextCase:   1,
		pq:         make(outputQueue, 0, 1),
		inputLock:  inputLock,
		caseOutput: caseOutput,
		caseNotify: caseNotify,
		endNotify:  endNotify,
		Input:      input,
		Output:     output,
	}

	heap.Init(&manager.pq)
	go manager.input()
	return manager
}

// NewGCJManager start the goroutine responsible by manage the io based on case
func NewGCJManager(caseTask GCJCase) *GCJManager {
	var err interface{}
	log.Init(log.DefaultLevel)

	var input *os.File
	if len(OptInput) > 0 {
		log.Debug.Printf("Opening input: %s\n", OptInput)
		input, err = os.Open(OptInput)
		if err != nil {
			panic(err)
		}
		//defer input.Close()
	} else {
		log.Debug.Println("Using stdin")
		input = os.Stdin
	}

	var output *os.File
	if len(OptOutput) > 0 {
		log.Debug.Printf("Opening output: %s\n", OptOutput)
		output, err = os.Create(OptOutput)
		if err != nil {
			panic(err)
		}
		//defer output.Close()
	} else {
		log.Debug.Println("Using stdout")
		output = os.Stdout
	}

	return NewGCJManagerIO(caseTask, input, output)
}

func (manager GCJManager) end() bool {
	return manager.nextCase > manager.cases
}

// Print send a request to print the return of a case
func (manager *GCJManager) Print(caseNum uint64, value string) {
	item := &googleCJCase{
		value:   value,
		caseNum: caseNum,
	}
	manager.caseOutput <- item
}

// WaitCaseEnd wait the end of one case and return the number of this case
func (manager *GCJManager) waitCaseEnd() uint64 {
	return <-manager.caseNotify
}

// WaitEnd wait the end of execution of all cases
func (manager *GCJManager) WaitEnd() {
	<-manager.endNotify
}

// InputUnlock free input to the next execution
func (manager *GCJManager) InputUnlock() {
	manager.inputLock <- true
}

// InputLock free input to the next execution
func (manager *GCJManager) InputLock() {
	<-manager.inputLock
}

func (manager GCJManager) notifyCaseEnd(caseNum uint64) {
	select {
	case manager.caseNotify <- caseNum:
	default:
	}
}

func (manager GCJManager) printCase(item *googleCJCase) {
	fmt.Fprintf(manager.Output, "Case #%d: %s\n", item.caseNum, item.value)
}

func (manager *GCJManager) input() {

	fmt.Fscanf(manager.Input, "%d", &manager.cases)

	go manager.output()

	for i := uint64(1); i <= manager.cases; i++ {

		log.Debug.Printf("Case %d lauching\n", i)
		go func(caseNum uint64, caseManager *GCJManager) {
			value := caseManager.caseTask(manager)
			log.Debug.Printf("Case %d returned\n", caseNum)
			manager.Print(caseNum, value)
		}(i, manager)

		manager.InputLock()
		log.Debug.Printf("Case %d input ended\n", i)
		if seqMode {
			ret := <-manager.caseNotify
			log.Info.Printf("Case %d ended\n", ret)
			if ret != i {
				panic("Notified another case. This is sequential mode!")
			}
		}
	}

	if manager.Input != os.Stdin {
		var input interface{} = manager.Input
		switch v := input.(type) {
		case *os.File:
			v.Close()
		}
	}
}

func (manager *GCJManager) output() {

	for !manager.end() {
		item := <-manager.caseOutput
		log.Debug.Printf("Case %d output received\n", item.caseNum)

		if item.caseNum != manager.nextCase {
			heap.Push(&manager.pq, item)
			log.Debug.Printf("Case %d output pushed\n", item.caseNum)
			continue
		}

		manager.printCase(item)
		manager.nextCase++
		manager.notifyCaseEnd(item.caseNum)

		for manager.pq.Len() > 0 && manager.nextCase == manager.pq.Head().caseNum {
			item := heap.Pop(&manager.pq).(*googleCJCase)
			manager.printCase(item)
			manager.nextCase++
			manager.notifyCaseEnd(item.caseNum)
		}
	}
	manager.endNotify <- true
	if manager.Output != os.Stdout {
		var output interface{} = manager.Output
		switch v := output.(type) {
		case *os.File:
			v.Close()
		}
	}
}
