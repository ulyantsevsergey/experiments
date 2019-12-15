package main

import (
	"awesomeProject/pool"
	work "awesomeProject/work"
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(WorkerCount) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, job := range work.CreateJobs(20) {
			work.DoWork(job, 1)
		}
	}
}
