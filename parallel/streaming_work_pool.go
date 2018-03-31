package parallel

import (
	"runtime"
)

type WorkItem interface{}

type workItemCompletionMsg struct {
	index int64
	w     WorkItem
}

type StreamingWorkPool struct {
	worker             func(i int64) (item WorkItem, err error)
	totalWorkSize      int64
	workItemRequest    chan int64
	workItemResponse   chan WorkItem
	responseBufferSize int64
	maxCacheSize       int64
}

// must only be called from a single thread!
func (p *StreamingWorkPool) Get(idx int64) WorkItem {
	p.workItemRequest <- idx
	return <-p.workItemResponse
}

func (p *StreamingWorkPool) workItemWorker(idx int64, response chan<- workItemCompletionMsg) {
	w, err := p.worker(idx)
	if err != nil {
		panic(err)
	} else {
		response <- workItemCompletionMsg{idx, w}
	}
}

func (p *StreamingWorkPool) workerManager() {
	unsortedResponses := make(chan workItemCompletionMsg, p.responseBufferSize)

	nextIndex := int64(0)
	//maxIndex := p.outWidth * p.outHeight
	maxIndex := p.totalWorkSize
	numRunningWorkers := 0
	maxNumRunningWorkers := runtime.GOMAXPROCS(-1)

	// start initial workers
	for i := int64(0); i < nextIndex; i++ {
		//go DownsamplePixel(ppm, nextIndex, factor, unsortedResponses)
		go p.workItemWorker(nextIndex, unsortedResponses)
		numRunningWorkers++
		nextIndex++
	}

	cache := make(map[int64]WorkItem)
	waitingFor := int64(-1)
	for {
		select {

		case reqIdx := <-p.workItemRequest:
			if px, ok := cache[reqIdx]; ok {
				delete(cache, reqIdx)
				p.workItemResponse <- px
			} else {
				waitingFor = reqIdx
				// start a new runner for this, possibly creating too many runners and possibly doing extra work
				//go DownsamplePixel(ppm, reqIdx, factor, unsortedResponses)
				go p.workItemWorker(reqIdx, unsortedResponses)
				numRunningWorkers++
				nextIndex = reqIdx + 1
			}

		case pxMsg := <-unsortedResponses:
			if pxMsg.index == waitingFor {
				// if we're waiting for it, send it directly
				p.workItemResponse <- pxMsg.w
				waitingFor = -1 // probably overkill
			} else {
				cache[pxMsg.index] = pxMsg.w
			}
			numRunningWorkers--

		}

		// start more workers if necessary
		if int64(len(cache)) < p.maxCacheSize && numRunningWorkers < maxNumRunningWorkers && nextIndex < maxIndex {
			for i := 0; i < (maxNumRunningWorkers - numRunningWorkers); i++ {
				//go DownsamplePixel(ppm, nextIndex, factor, unsortedResponses)
				go p.workItemWorker(nextIndex, unsortedResponses)
				numRunningWorkers++
				nextIndex++
			}
		}

	}
}

func MakeStreamingWorkPool(totalWorkSize int64, worker func(i int64) (item WorkItem, err error)) *StreamingWorkPool {
	out := StreamingWorkPool{
		totalWorkSize:      totalWorkSize,
		worker:             worker,
		workItemRequest:    make(chan int64),
		workItemResponse:   make(chan WorkItem),
		responseBufferSize: 10240,
		maxCacheSize:       10240,
	}

	go out.workerManager()

	return &out
}
