package manager

import (
	"bufio"
	"github.com/tiagofalcao/GoNotebook/log"
	"io"
	"os"
)

type Manager struct {
	Input      *bufio.Reader
	inputFile  *os.File
	Output     io.Writer
	outputFile *os.File
}

type Task func(r *bufio.Reader, w io.Writer)

// NewManagerFile start the goroutine responsible by manage the io based on case
func NewManagerFile(task Task, input *bufio.Reader, inputFile *os.File, output io.Writer, outputFile *os.File) {
	log.Init(log.DefaultLevel)

	manager := &Manager{
		Input:      input,
		inputFile:  inputFile,
		Output:     output,
		outputFile: outputFile,
	}
	defer manager.Close()

	task(input, output)
}

// NewManagerIO start the goroutine responsible by manage the io based on case
func NewManagerIO(task Task, input *bufio.Reader, output io.Writer) {
	NewManagerFile(task, input, nil, output, nil)
}

// NewManager start the goroutine responsible by manage the io based on case
func NewManager(task Task) {
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
		if FlushMode {
			output = outputFile
		} else {
			output = bufio.NewWriter(outputFile)
		}
	} else {
		log.Debug.Println("Using stdout")
		if FlushMode {
			output = os.Stdout
		} else {
			output = bufio.NewWriter(os.Stdout)
		}
	}

	NewManagerFile(task, input, inputFile, output, outputFile)
}

func (manager *Manager) Close() {
	var output interface{} = manager.Output
	switch v := output.(type) {
	case bufio.Writer:
		v.Flush()
	}

	if manager.outputFile != nil {
		manager.outputFile.Close()
	}
}
