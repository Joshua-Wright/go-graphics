package parallel

import (
	"sync"
	"runtime"
	"time"
)

func parallelForWorker(wg *sync.WaitGroup, jobs chan int, f func(int)) {
	for i := range jobs {
		f(i)
		wg.Done()
	}
}

func ParallelFor(start, end int, f func(int)) {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go parallelForWorker(&wg, jobs, f)
	}

	// queue
	wg.Add(end - start)
	for i := start; i < end; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func ParallelFuncs(funcs ...func()) {
	ParallelFor(0, len(funcs), func(i int) {
		funcs[i]()
	})
}

/////////////////////////////////////////////////////

type jobRange struct {
	startInclusive, endExclusive int
}

type timingResult struct {
	elapsed time.Duration
	size    int
}

func parallelForAdaptiveWorker(wg *sync.WaitGroup, jobs chan jobRange, results chan timingResult, f func(int, int)) {
	for r := range jobs {
		start := time.Now()
		f(r.startInclusive, r.endExclusive)
		elapsedTime := time.Since(start)
		wg.Done()
		results <- timingResult{
			elapsedTime,
			r.endExclusive - r.startInclusive,
		}
	}
}

const targetJobRangeTime time.Duration = time.Millisecond * 100

func ParallelForAdaptive(start, end int, f func(startInclusive, endExclusive int)) (meanJobTime time.Duration) {
	numCPUs := runtime.GOMAXPROCS(-1)
	if numCPUs >= (end - start) {
		// TODO fallback case for when numCPUs <= (end-start)
		println(numCPUs, start, end)
		panic("too few jobs")
	}
	jobRanges := make(chan jobRange)
	timingResults := make(chan timingResult)
	var wg sync.WaitGroup

	work_completed := 0
	totalWorkTime := time.Duration(0)

	// start workers
	for i := 0; i < numCPUs; i++ {
		go parallelForAdaptiveWorker(&wg, jobRanges, timingResults, f)
	}

	// queue initial jobs to gather timing information
	wg.Add(numCPUs)
	for i := start; i < start+numCPUs; i++ {
		jobRanges <- jobRange{i, i + 1}
	}

	nextJob := numCPUs
	for {
		timing := <-timingResults
		if (nextJob == end) {
			break
		}
		work_completed += timing.size
		totalWorkTime += timing.elapsed
		durationPerJob := int64(totalWorkTime) / int64(work_completed)
		//println("duration per job:", durationPerJob)
		nextJobSize := int(int64(targetJobRangeTime) / durationPerJob)
		if nextJob+nextJobSize > end {
			nextJobSize = end - nextJob
		}
		wg.Add(1)
		jobRanges <- jobRange{nextJob, nextJob + nextJobSize}
		nextJob += nextJobSize
	}

	close(jobRanges)
	wg.Wait()
	meanJobTime = totalWorkTime / time.Duration(end-start)
	return meanJobTime
}
