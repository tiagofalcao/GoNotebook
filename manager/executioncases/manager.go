package executioncases

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/log"
	"io"
	"os"
)

type ExecutionManager struct {
	cases      uint64         // Amount of Cases
	caseTask   ExecutionCase  // Execute Case function
	nextCase   uint64         // The Next Case
	casePrint  CasePrint      // Print Case answer
	pq         outputQueue
	inputLock  chan bool
	caseOutput chan *caseOutput
	caseNotify chan uint64
	endNotify  chan bool
	Input      *bufio.Reader
	inputFile  *os.File
	Output     io.Writer
	outputFile *os.File
}

// NewExecutionManagerFile start the goroutine responsible by manage the io based on case
func NewExecutionManagerFile(caseTask ExecutionCase, print CasePrint, input *bufio.Reader, inputFile *os.File, output io.Writer, outputFile *os.File) *ExecutionManager {
	log.Init(log.DefaultLevel)

	caseOutput := make(chan *caseOutput, 100)
	caseNotify := make(chan uint64, 10)
	endNotify := make(chan bool)
	inputLock := make(chan bool)

	manager := &ExecutionManager{
		caseTask:   caseTask,
		nextCase:   1,
		casePrint:  print,
		pq:         make(outputQueue, 0, 1),
		inputLock:  inputLock,
		caseOutput: caseOutput,
		caseNotify: caseNotify,
		endNotify:  endNotify,
		Input:      input,
		inputFile:  inputFile,
		Output:     output,
		outputFile: outputFile,
	}

	heap.Init(&manager.pq)
	go manager.input()
	return manager
}

// NewExecutionManagerIO start the goroutine responsible by manage the io based on case
func NewExecutionManagerIO(caseTask ExecutionCase, print CasePrint, input *bufio.Reader, output io.Writer) *ExecutionManager {
	return NewExecutionManagerFile(caseTask, print, input, nil, output, nil)
}

// NewExecutionManager start the goroutine responsible by manage the io based on case
func NewExecutionManager(caseTask ExecutionCase, print CasePrint) *ExecutionManager {
	var err interface{}
	log.Init(log.DefaultLevel)

	var input *bufio.Reader
	var output io.Writer

	var inputFile *os.File
	if len(OptInput) > 0 {
		log.Debug.Printf("Opening input: %s\n", OptInput)
		inputFile, err = os.Open(OptInput)
		if err != nil {
			panic(err)
		}
		input = bufio.NewReader(inputFile)
	} else {
		log.Debug.Println("Using stdin")
		input = bufio.NewReader(os.Stdin)
	}

	var outputFile *os.File
	if len(OptOutput) > 0 {
		log.Debug.Printf("Opening output: %s\n", OptOutput)
		outputFile, err = os.Create(OptOutput)
		if err != nil {
			panic(err)
		}
		if flushMode {
			output = outputFile
		} else {
			output = bufio.NewWriter(outputFile)
		}
	} else {
		log.Debug.Println("Using stdout")
		if flushMode {
			output = os.Stdout
		} else {
			output = bufio.NewWriter(os.Stdout)
		}
	}

	return NewExecutionManagerFile(caseTask, print, input, inputFile, output, outputFile)
}

func (manager ExecutionManager) end() bool {
	return manager.nextCase > manager.cases
}

// Print send a request to print the return of a case
func (manager *ExecutionManager) Print(caseNum uint64, value string) {
	item := &caseOutput{
		value:   value,
		caseNum: caseNum,
	}
	manager.caseOutput <- item
}

// WaitCaseEnd wait the end of one case and return the number of this case
func (manager *ExecutionManager) waitCaseEnd() uint64 {
	return <-manager.caseNotify
}

// WaitEnd wait the end of execution of all cases
func (manager *ExecutionManager) WaitEnd() {
	<-manager.endNotify
}

// InputUnlock free input to the next execution
func (manager *ExecutionManager) InputUnlock() {
	manager.inputLock <- true
}

// InputLock free input to the next execution
func (manager *ExecutionManager) InputLock() {
	<-manager.inputLock
}

func (manager ExecutionManager) notifyCaseEnd(caseNum uint64) {
	select {
	case manager.caseNotify <- caseNum:
	default:
	}
}

func (manager *ExecutionManager) input() {

	log.Debug.Println(manager.Input.Peek(10))
	fmt.Fscanf(manager.Input, "%d\n", &manager.cases)
	log.Debug.Println(manager.Input.Peek(10))
	log.Debug.Printf("Running %d cases", manager.cases)

	go manager.output()

	for i := uint64(1); i <= manager.cases; i++ {

		log.Debug.Printf("Case %d lauching\n", i)
		go func(caseNum uint64, caseManager *ExecutionManager) {
			value := caseManager.caseTask(manager.Input, manager.inputLock)
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

	if manager.inputFile != nil {
		manager.inputFile.Close()
	}
}

func (manager *ExecutionManager) output() {

	for !manager.end() {
		item := <-manager.caseOutput
		log.Debug.Printf("Case %d output received\n", item.caseNum)

		if item.caseNum != manager.nextCase {
			heap.Push(&manager.pq, item)
			log.Debug.Printf("Case %d output pushed\n", item.caseNum)
			continue
		}

		manager.casePrint(manager.Output, item.caseNum, item.value)
		manager.nextCase++
		manager.notifyCaseEnd(item.caseNum)

		for manager.pq.Len() > 0 && manager.nextCase == manager.pq.Head().caseNum {
			item := heap.Pop(&manager.pq).(*caseOutput)
			manager.casePrint(manager.Output, item.caseNum, item.value)
			manager.nextCase++
			manager.notifyCaseEnd(item.caseNum)
		}
	}
	manager.endNotify <- true

	var output interface{} = manager.Output
	switch v := output.(type) {
	case bufio.Writer:
		v.Flush()
	}

	if manager.outputFile != nil {
		manager.outputFile.Close()
	}

}
