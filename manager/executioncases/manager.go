package executioncases

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/log"
	"io"
	"os"
	"runtime"
)

import coremanager "github.com/tiagofalcao/GoNotebook/manager"

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
	Timing     []timing
}

// NewExecutionManagerFile start the goroutine responsible by manage the io based on case
func NewExecutionManagerFile(caseTask ExecutionCase, print CasePrint, input *bufio.Reader, inputFile *os.File, output io.Writer, outputFile *os.File) *ExecutionManager {
	log.Init(log.DefaultLevel)

	caseOutput := make(chan *caseOutput, 100)
	caseNotify := make(chan uint64, 100)
	endNotify := make(chan bool)
	inputLock := make(chan bool)

	manager := &ExecutionManager{
		caseTask:   caseTask,
		nextCase:   0,
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
		Timing:     nil,
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
	if len(coremanager.OptInput) > 0 {
		log.Debug.Printf("Opening input: %s\n", coremanager.OptInput)
		inputFile, err = os.Open(coremanager.OptInput)
		if err != nil {
			panic(err)
		}
		input = bufio.NewReader(inputFile)
	} else {
		log.Debug.Println("Using stdin")
		input = bufio.NewReader(os.Stdin)
	}

	var outputFile *os.File
	if len(coremanager.OptOutput) > 0 {
		log.Debug.Printf("Opening output: %s\n", coremanager.OptOutput)
		outputFile, err = os.Create(coremanager.OptOutput)
		if err != nil {
			panic(err)
		}
		if coremanager.FlushMode {
			output = outputFile
		} else {
			output = bufio.NewWriter(outputFile)
		}
	} else {
		log.Debug.Println("Using stdout")
		if coremanager.FlushMode {
			output = os.Stdout
		} else {
			output = bufio.NewWriter(os.Stdout)
		}
	}

	return NewExecutionManagerFile(caseTask, print, input, inputFile, output, outputFile)
}

func (manager ExecutionManager) end() bool {
	return manager.nextCase >= manager.cases
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
	manager.caseNotify <- caseNum
}

func (manager ExecutionManager) caseEnd(caseNum, running uint64) uint64 {
	t := manager.Timing[caseNum]
	complete := t.complete()
	reading := t.reading()
	computing := t.reading()
	log.Info.Printf("Case %d ended: %v [%v | %v]\n", caseNum, complete, reading, computing )
	if computing > (2 * reading) && ParallelMode < runtime.NumCPU() {
		ParallelMode += 1
	}
	return running - 1
}

func (manager *ExecutionManager) input() {

	fmt.Fscanf(manager.Input, "%d\n", &manager.cases)
	log.Debug.Printf("Running %d cases", manager.cases)

	manager.Timing = make([]timing, manager.cases)

	go manager.output()

	running := uint64(0)

	for i := uint64(0); i < manager.cases; i++ {

		log.Debug.Printf("Case %d lauching\n", i)
		running += 1

		manager.Timing[i].start()
		go func(caseNum uint64, caseManager *ExecutionManager) {
			value := caseManager.caseTask(manager.Input, manager.inputLock)
			log.Debug.Printf("Case %d returned\n", caseNum)
			manager.Print(caseNum, value)
		}(i, manager)

		manager.InputLock()
    manager.Timing[i].read()
		log.Debug.Printf("Case %d input ended\n", i)

		log.Info.Printf("Running %d cases\n", running)

		// Try fetch a case notification
		select {
		case ret := <-manager.caseNotify:
		running = manager.caseEnd(ret, running)
		default:
		}

		// Too many running, wait to fetch a case notification
		for running >= uint64(ParallelMode) {
			ret := <-manager.caseNotify
			running = manager.caseEnd(ret, running)
		}
	}

	if manager.inputFile != nil {
		manager.inputFile.Close()
	}
}

func (manager *ExecutionManager) output() {

	for !manager.end() {
		item := <-manager.caseOutput
		manager.Timing[item.caseNum].end()
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
