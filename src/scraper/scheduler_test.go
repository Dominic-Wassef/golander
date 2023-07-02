package scraper

import (
	"testing"
	"time"
)

func TestTaskAndScheduler(t *testing.T) {
	taskFunc := func() error {
		return nil
	}

	errorHandler := func(error) {}

	task := NewTask(time.Millisecond*100, taskFunc, errorHandler)

	if task.Running {
		t.Errorf("newly created task should not be running")
	}

	scheduler := NewScheduler()

	scheduler.AddTask(task)

	scheduler.StartAll()

	if !task.Running {
		t.Errorf("task should be running after StartAll")
	}

	scheduler.StopAll()

	if task.Running {
		t.Errorf("task should not be running after StopAll")
	}
}
