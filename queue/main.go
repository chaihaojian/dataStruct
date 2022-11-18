package main

import (
	"DataStruct/queue/circlequeue"
	"DataStruct/queue/singlequeue"
	"fmt"
	"os"
)

func main() {
	var class string

	fmt.Println("==========================")
	fmt.Println("chose queue class:")
	fmt.Scan(&class)
	switch class {
	case "single":
		queue := &singlequeue.SingleQueue{
			MaxSize: 5,
			Front:   -1,
			Rear:    -1,
		}
		SingleOperation(queue)
	case "circle":
		queue := &circlequeue.CircleQueue{
			MaxSize: 5,
			Front:   0,
			Rear:    0,
		}
		CircleOperation(queue)
	case "quit":
		os.Exit(0)
	}
}

func SingleOperation(queue *singlequeue.SingleQueue) {
	for {
		var key string
		var val int
		fmt.Println("==========================")
		fmt.Println("select operation:")
		fmt.Println("push")
		fmt.Println("pop")
		fmt.Println("show")
		fmt.Println("==========================")
		fmt.Scan(&key)
		switch key {
		case "push":
			fmt.Println("input:")
			fmt.Scan(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("push failed", err)
			}
			queue.Show()
		case "pop":
			v, err := queue.Pop()
			if err != nil {
				fmt.Println("pop failed", err)
			}
			fmt.Println(v)
			queue.Show()
		case "show":
			queue.Show()
		case "quit":
			os.Exit(0)
		}
	}
}

func CircleOperation(queue *circlequeue.CircleQueue) {
	for {
		var key string
		var val int
		fmt.Println("==========================")
		fmt.Println("select operation:")
		fmt.Println("push")
		fmt.Println("pop")
		fmt.Println("capacity")
		fmt.Println("show")
		fmt.Println("==========================")
		fmt.Scan(&key)
		switch key {
		case "push":
			fmt.Println("input:")
			fmt.Scan(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("push failed", err)
			}
			queue.Show()
		case "pop":
			v, err := queue.Pop()
			if err != nil {
				fmt.Println("pop failed", err)
			}
			fmt.Println(v)
			queue.Show()
		case "capacity":
			fmt.Println(queue.Cap())
		case "show":
			queue.Show()
		case "quit":
			os.Exit(0)
		}
	}
}
