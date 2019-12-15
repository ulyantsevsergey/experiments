package pool

import (
	work "awesomeProject/work"
	"log"
)

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	Channel       chan Work
	End           chan bool
}

func createWorker(id int, WorkerChannel chan chan Work) Worker {
	log.Println("starting worker: ", id)
	Worker := Worker{
		ID:            id,
		Channel:       make(chan Work),
		WorkerChannel: WorkerChannel,
		End:           make(chan bool),
	}
	return Worker
}

// start worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case w := <-w.Channel: // worker has received job
				work.DoWork(w.Job, w.ID) // do work
			case <-w.End:
				return
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}
