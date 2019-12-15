package job

import (
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// create random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Job() string {
	job := RandStringRunes(8)
	return job
}

// create list of jobs
func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, Job())
	}
	return jobs
}

// mimics any type of job that can be run concurrently
func DoWork(word string, id int) {
	h := fnv.New32a()
	_, err := h.Write([]byte(word))
	time.Sleep(time.Second)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)
	}
}
