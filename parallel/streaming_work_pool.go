package parallel

import (
	"runtime"
)

type WorkItem interface{}

type workItemChunk struct {
	startIncl, endExcl int64
	items              []WorkItem
}

type StreamingWorkPool struct {
	worker             func(i int64) (item WorkItem, err error)
	totalWorkSize      int64
	responseBufferSize int64
	maxCacheSize       int64
	chunkSize          int64

	currentWorkCache workItemChunk
	workItemRequest  chan int64
	workItemResponse chan workItemChunk
}

// must only be called from a single thread!
func (p *StreamingWorkPool) Get(idx int64) WorkItem {
	if idx < p.currentWorkCache.startIncl || idx >= p.currentWorkCache.endExcl {
		p.workItemRequest <- idx
		p.currentWorkCache = <-p.workItemResponse
	}
	return p.currentWorkCache.items[idx-p.currentWorkCache.startIncl]
}

func (p *StreamingWorkPool) workItemWorker(startIncl int64, response chan<- workItemChunk) {
	out := workItemChunk{
		startIncl: startIncl,
		endExcl:   startIncl + p.chunkSize,
		items:     make([]WorkItem, p.chunkSize),
	}
	for i := out.startIncl; i < out.endExcl; i++ {
		w, err := p.worker(i)
		if err != nil {
			panic(err)
		}
		out.items[i-out.startIncl] = w
	}
	response <- out
}

func (p *StreamingWorkPool) workerManager() {
	unsortedResponses := make(chan workItemChunk, p.responseBufferSize)

	nextIndex := int64(0)
	maxIndex := p.totalWorkSize
	numRunningWorkers := 0
	maxNumRunningWorkers := runtime.GOMAXPROCS(-1)

	// start initial workers
	for i := int64(0); i < nextIndex; i++ {
		go p.workItemWorker(nextIndex, unsortedResponses)
		numRunningWorkers++
		nextIndex++
	}

	// cache maps starting index to work item chunk
	cache := make(map[int64]workItemChunk, p.maxCacheSize)
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
				go p.workItemWorker(reqIdx, unsortedResponses)
				numRunningWorkers++
				nextIndex = reqIdx + p.chunkSize
			}

		case pxMsg := <-unsortedResponses:
			if pxMsg.startIncl == waitingFor {
				// if we're waiting for it, send it directly
				p.workItemResponse <- pxMsg
				waitingFor = -1 // probably overkill
			} else {
				cache[pxMsg.startIncl] = pxMsg
			}
			numRunningWorkers--

		}

		if int64(len(cache)) == p.maxCacheSize && numRunningWorkers < maxNumRunningWorkers/2 {
			// clear the cache if we're getting bogged down
			cache = make(map[int64]workItemChunk, p.maxCacheSize)
		}

		// start more workers if necessary
		if int64(len(cache)) < p.maxCacheSize && numRunningWorkers < maxNumRunningWorkers && nextIndex < maxIndex {
			for i := 0; i < (maxNumRunningWorkers - numRunningWorkers); i++ {
				go p.workItemWorker(nextIndex, unsortedResponses)
				numRunningWorkers++
				nextIndex += p.chunkSize
			}
		}

	}
}

func MakeStreamingWorkPool0(
	totalWorkSize, responseBufferSize, maxCacheSize, chunkSize int64,
	worker func(i int64) (item WorkItem, err error)) *StreamingWorkPool {

	out := StreamingWorkPool{
		totalWorkSize:      totalWorkSize,
		worker:             worker,
		workItemRequest:    make(chan int64),
		workItemResponse:   make(chan workItemChunk),
		responseBufferSize: responseBufferSize,
		maxCacheSize:       maxCacheSize,
		chunkSize:          chunkSize,
		currentWorkCache: workItemChunk{
			startIncl: -1,
			endExcl:   -1,
		},
	}

	go out.workerManager()

	return &out
}

func MakeStreamingWorkPool(totalWorkSize int64, worker func(i int64) (item WorkItem, err error)) *StreamingWorkPool {
	out := StreamingWorkPool{
		totalWorkSize:      totalWorkSize,
		worker:             worker,
		workItemRequest:    make(chan int64),
		workItemResponse:   make(chan workItemChunk),
		responseBufferSize: 10240,
		maxCacheSize:       10240,
		chunkSize:          128,
		currentWorkCache: workItemChunk{
			startIncl: -1,
			endExcl:   -1,
		},
	}

	go out.workerManager()

	return &out
}
