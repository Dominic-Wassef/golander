package scraper

import (
	"fmt"
	"sync"
	"time"
)

type TaskFunc func() error

type Task struct {
	Interval     time.Duration
	LastRun      time.Time
	Running      bool
	stop         chan bool
	taskFunc     TaskFunc
	errorHandler func(error)
}

func NewTask(interval time.Duration, taskFunc TaskFunc, errorHandler func(error)) *Task {
	return &Task{
		Interval:     interval,
		taskFunc:     taskFunc,
		errorHandler: errorHandler,
		stop:         make(chan bool),
	}
}

func (t *Task) Start() {
	t.Running = true

	go func() {
		ticker := time.NewTicker(t.Interval)

		for {
			select {
			case <-ticker.C:
				t.LastRun = time.Now()
				fmt.Printf("Starting task at %v\n", t.LastRun)
				if err := t.taskFunc(); err != nil {
					t.errorHandler(err)
				}
			case <-t.stop:
				ticker.Stop()
				return
			}
		}
	}()
}

func (t *Task) Stop() {
	t.Running = false
	t.stop <- true
}

type Scheduler struct {
	tasks []*Task
	mutex sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) AddTask(t *Task) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) StartAll() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, task := range s.tasks {
		if !task.Running {
			task.Start()
		}
	}
}

func (s *Scheduler) StopAll() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, task := range s.tasks {
		if task.Running {
			task.Stop()
		}
	}
}
