package circlequeue

import (
	"errors"
	"fmt"
)

// CircleQueue 数组实现循环队列
type CircleQueue struct {
	MaxSize int
	Front   int //0
	Rear    int //0
	Array   [5]int
}

func (this *CircleQueue) Push(val int) error {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.Array[this.Rear] = val
	this.Rear = (this.Rear + 1) % this.MaxSize
	return nil
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.Array[this.Front]
	this.Front = (this.Front + 1) % this.MaxSize
	return val, nil
}

func (this *CircleQueue) IsFull() bool {
	return (this.Rear+1)%this.MaxSize == this.Front
}

func (this *CircleQueue) IsEmpty() bool {
	return this.Front == this.Rear
}

func (this *CircleQueue) Cap() int {
	return (this.Rear - this.Front + this.MaxSize) % this.MaxSize
}

func (this *CircleQueue) Show() {
	temp := this.Front
	for i := 0; i < this.Cap(); i++ {
		fmt.Printf("%d -> ", this.Array[temp])
		temp = (temp + 1) % this.MaxSize
	}
	fmt.Println()
}
