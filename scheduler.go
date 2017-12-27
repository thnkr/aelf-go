package main

// ResType represents a Resource type
type ResType uint32

// Task defines a task in the graph with dependecies
type Task struct {
	f       func() // the transaction as a closure
	rts     ResType
	related []*Task
}

// TaskGroup defines group of related tasks in the graph
type TaskGroup struct {
	tasks     []Task
	resources map[ResType]bool // for fast querying of resource type
}

// Scheduler defines a batch of tasks to be executed
type Scheduler struct {
	txs map[ResType][]func() // collect resource usage
}

// NewScheduler creates a scheduler for batch processing
func NewScheduler() *Scheduler {
	s := &Scheduler{}
	s.txs = make(map[ResType][]func())
	return s
}

// AddTransaction will group dependent taskes together
func (s *Scheduler) AddTransaction(tx func(), rts []ResType) {
	for _, res := range rts {
		s.txs[res] = append(s.txs[res], tx)
	}
}
