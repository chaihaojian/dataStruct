package singlequeue

import (
	"errors"
	"fmt"
)

type SingleQueue struct {
	MaxSize int
	Front   int
	Rear    int
	Array   [10]int
}

func (this *SingleQueue) Push(val int) error {
	//检查队列是否已满
	if this.Rear == this.MaxSize-1 {
		return errors.New("queue full")
	}

	//push
	this.Rear++
	this.Array[this.Rear] = val
	return nil
}

func (this *SingleQueue) Pop() (int, error) {
	//检查队列是否为空
	if this.Front == this.Rear {
		return 0, errors.New("queue empty")
	}

	//pop
	this.Front++
	return this.Array[this.Front], nil
}

func (this *SingleQueue) Show() {
	for i := this.Front + 1; i <= this.Rear; i++ {
		fmt.Printf("%d -> ", this.Array[i])
	}
	fmt.Println()
}
