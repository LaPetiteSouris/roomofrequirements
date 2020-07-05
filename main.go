package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Worker is the work horse
type Worker interface {
	executeTask(string, *sync.WaitGroup) error
}

// VizWorker is just a type of worker
type VizWorker struct {
	id string
}

func (w *VizWorker) executeTask(task string, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Printf("Worker's id %s , executing task, message is %s \n", w.id, task)
	return nil
}

// WorkerPool struct
type WorkerPool struct {
	pool []Worker
	wg   *sync.WaitGroup
}

func (wp *WorkerPool) job(task string) error {
	// look into the pool
	// take out one worker goroutine
	worker := wp.pool[rand.Intn(len(wp.pool))]
	// increase waitgroup
	wp.wg.Add(1)
	// execute job
	go worker.executeTask(task, wp.wg)
	return nil
}

// Queue struct
type Queue struct {
	queue chan string
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	// Initiate worker pool
	numberOfWorkers := 3
	vizWorkerArray := make([]Worker, 0)
	for i := 0; i < numberOfWorkers; i++ {
		w := &VizWorker{id: strconv.FormatInt(int64(i), 10)}
		vizWorkerArray = append(vizWorkerArray, w)
	}

	workerPool := &WorkerPool{wg: &wg, pool: vizWorkerArray}

	messages := []string{"alpha", "beta", "gamma"}
	for _, v := range messages {
		workerPool.job(v)
	}

	wg.Wait()
	// Add to pool

	//queue := &Queue{queue: make(chan string)}

	// initiate worker pool

}
