package workerpool

import (
	"ai-service/internal/service"
	"fmt"
)

type Pool struct {
	WorkerCount int
	JobBuffer   int
	Jobs        chan []byte
	TestService *service.TestService
}

func New(workerCount, jobBuffer int, testService *service.TestService) *Pool {
	return &Pool{
		WorkerCount: workerCount,
		JobBuffer:   jobBuffer,
		Jobs:        make(chan []byte, jobBuffer),
		TestService: testService,
	}
}


func (p *Pool) Start() {
	for i := 1; i <= p.WorkerCount; i++ {
		go p.worker(i)
	}
	fmt.Printf(" %d workers ishga tushdi\n", p.WorkerCount)
}


func (p *Pool) worker(id int) {
	fmt.Printf("Worker %d ready\n", id)
	for job := range p.Jobs {
		service.ProcessMessage(job, p.TestService)
	}
}


func (p *Pool) Submit(job []byte) {
	p.Jobs <- job
}
