package executioncases

type outputQueue []*caseOutput

func (output outputQueue) Len() int { return len(output) }

func (output outputQueue) Less(i, j int) bool {
	return output[i].caseNum < output[j].caseNum
}

func (output outputQueue) Swap(i, j int) {
	output[i], output[j] = output[j], output[i]
	output[i].index = i
	output[j].index = j
}

func (output *outputQueue) Push(x interface{}) {
	n := output.Len()
	item := x.(*caseOutput)
	item.index = n
	*output = append(*output, item)
}

func (output outputQueue) Head() *caseOutput {
	return output[0]
}

func (output *outputQueue) Pop() interface{} {
	old := *output
	n := len(old)
	item := old[n-1]
	item.index = -1
	*output = old[0 : n-1]
	return item
}
