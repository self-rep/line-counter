package worker

import "sync"

type QueueWorker struct {
	addReq   chan int
	delReq   chan int
	resetReq chan int
}

var (
	Safe        = sync.RWMutex{}
	GlobalLines = 0
)

func StartQueue() *QueueWorker {
	c := &QueueWorker{addReq: make(chan int, 128), delReq: make(chan int, 128), resetReq: make(chan int, 128)}
	go c.worker()
	return c
}

func (w *QueueWorker) Iterate() {
	w.addReq <- 1337 // send something down channel to trigger
}

func (w *QueueWorker) Deiterate() {
	w.delReq <- 1337
}

func (w *QueueWorker) Reset() {
	w.resetReq <- 1337
}

//worker is needed to safley iterate
func (w *QueueWorker) worker() {
	for {
		select {
		case <-w.addReq:
			Safe.Lock()
			GlobalLines++
			Safe.Unlock()
		case <-w.delReq:
			Safe.Lock()
			GlobalLines--
			Safe.Unlock()
		case <-w.resetReq:
			Safe.Lock()
			GlobalLines = 0
			Safe.Unlock()
		}
	}
}
