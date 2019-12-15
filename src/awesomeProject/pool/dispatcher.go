package pool

import "log"

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work chan Work // receives jobs to send to workers
	End  chan bool // when receives bool stops workers
}

func StartDispatcher(workerCount int) Collector {
	var workers []Worker
	input := make(chan Work) // channel to receive work
	end := make(chan bool)   // channel to spin down workers
	collector := Collector{Work: input, End: end}

	for i := 1; i <= workerCount; i++ {
		worker := createWorker(i, WorkerChannel)
		worker.Start()
		workers = append(workers, worker) // stores worker
	}
	// start collector
	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					log.Println("stopping worker: ", w.ID)
					w.Stop() // stop worker
				}
				return
			case work := <-input:
				log.Println("processing job: ", work.ID)
				worker := <-WorkerChannel // wait for available channel
				worker <- work            // dispatch work to worker
			}
		}
	}()

	return collector
}
