package manager

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/log"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
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
	Input      *bufio.Reader
	inputFile  *os.File
	Output     io.Writer
	outputFile *os.File
}

// NewGCJManagerFile start the goroutine responsible by manage the io based on case
func NewGCJManagerFile(caseTask GCJCase, input *bufio.Reader, inputFile *os.File, output io.Writer, outputFile *os.File) *GCJManager {
	log.Init(log.DefaultLevel)

	caseOutput := make(chan *googleCJCase, 100)
	caseNotify := make(chan uint64, 10)
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
		inputFile:  inputFile,
		Output:     output,
		outputFile: outputFile,
	}

	heap.Init(&manager.pq)
	go manager.input()
	return manager
}

// NewGCJManagerIO start the goroutine responsible by manage the io based on case
func NewGCJManagerIO(caseTask GCJCase, input *bufio.Reader, output io.Writer) *GCJManager {
	return NewGCJManagerFile(caseTask, input, nil, output, nil)
}

// NewGCJManager start the goroutine responsible by manage the io based on case
func NewGCJManager(caseTask GCJCase) *GCJManager {
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
		output = bufio.NewWriter(outputFile)
	} else {
		log.Debug.Println("Using stdout")
		output = bufio.NewWriter(os.Stdout)
	}

	return NewGCJManagerFile(caseTask, input, inputFile, output, outputFile)
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
	first, size := utf8.DecodeRuneInString(item.value)
	if size == 0 {
		log.Error.Println("Non utf8 rune")
	}
	if unicode.IsSpace(first) {
		fmt.Fprintf(manager.Output, "Case #%d:%s\n", item.caseNum, item.value)
		return
	}
	fmt.Fprintf(manager.Output, "Case #%d: %s\n", item.caseNum, item.value)
}

func (manager *GCJManager) input() {

	log.Debug.Println(manager.Input.Peek(10))
	fmt.Fscanf(manager.Input, "%d\n", &manager.cases)
	log.Debug.Println(manager.Input.Peek(10))
	log.Debug.Printf("Running %d cases", manager.cases)

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

	if manager.inputFile != nil {
		manager.inputFile.Close()
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

	var output interface{} = manager.Output
	switch v := output.(type) {
	case bufio.Writer:
		v.Flush()
	}

	if manager.outputFile != nil {
		manager.outputFile.Close()
	}

}
