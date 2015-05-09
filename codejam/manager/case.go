package manager

type GCJCase func(caseManager *GCJManager) string

type googleCJCase struct {
	value   string // The result
	caseNum uint64 // The case number
	index   int    // The index of the item in the heap.
}
