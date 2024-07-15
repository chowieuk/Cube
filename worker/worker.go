package worker

import (
	"fmt"

	"github.com/chowieuk/Cube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectState() {
	fmt.Println("I will collect state")
}
func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}
func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}
func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
