package main

import (
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
		Operation(queue)
	case "quit":
		os.Exit(0)
	}
}

func Operation(queue *singlequeue.SingleQueue) {
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
