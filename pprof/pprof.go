package pprof

import (
	"fmt"
	"math"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	heapStore   []string
	leakedChans []chan struct{}
)

func Pprof() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	http.HandleFunc("/cpu", cpuSpike)
	http.HandleFunc("/heap", heapSpike)
	http.HandleFunc("/alloc", allocSpike)
	http.HandleFunc("/mutex", mutexSpike)
	http.HandleFunc("/block", blockSpike)
	http.HandleFunc("/goleak", goRoutineLeakSpike)
	http.HandleFunc("/clear", clearLeakHandler)

	fmt.Println("pprof server running on :8080")
	fmt.Println("Try: http://localhost:8080/debug/pprof")
	http.ListenAndServe(":8080", nil)
}

func cpuSpike(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	sum := 0.0

	for i := 0; i < 80_000_000; i++ {
		sum += math.Sin(float64(i)) * math.Cos(float64(i))
	}

	fmt.Fprintf(w, "CPU spike done in %v\n", time.Since(start))
}

func heapSpike(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for i := 0; i < 30; i++ {
		s := strings.Repeat("X", 1_000_000)
		heapStore = append(heapStore, s)
	}

	fmt.Fprintf(w, "Heap increased by ~30MB in %v\n", time.Since(start))
}

func allocSpike(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for i := 0; i < 5_000_000; i++ {
		_ = fmt.Sprintf("alloc-%d", i)
	}

	fmt.Fprintf(w, "Alloc spike (5M allocations) done in %v\n", time.Since(start))
}

func mutexSpike(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	start := time.Now()
	workers := 100

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(5 * time.Millisecond)
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Fprintf(w, "Mutex contention test done in %v\n", time.Since(start))
}

func blockSpike(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 50; i++ {
			<-ch
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val
		}(i)
	}

	wg.Wait()
	fmt.Fprintf(w, "Blocking simulation complete (slow consumer)\n")
}

func goRoutineLeakSpike(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 500; i++ {
		ch := make(chan struct{})
		leakedChans = append(leakedChans, ch)

		go func(c chan struct{}) {
			<-c
		}(ch)
	}
	fmt.Fprintf(w, "Leaked 500 goroutines\n")
}

func clearLeakHandler(w http.ResponseWriter, r *http.Request) {
	for _, ch := range leakedChans {
		close(ch)
	}
	leakedChans = nil
	fmt.Fprintf(w, "Cleared leaked goroutines\n")
}
