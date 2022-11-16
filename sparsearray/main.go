package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//初始化一个二维数组,以切片的方式
	row := 6 //行
	col := 6 //列
	val := 0 //默认值

	Arr := make([][]int, row)
	for i := range Arr {
		Arr[i] = make([]int, col)
	}

	Arr[1][2] = 1
	Arr[3][2] = 1
	Arr[1][4] = 2
	Arr[5][2] = 2

	for _, r := range Arr {
		for _, v := range r {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}

	//构建稀疏数组
	type ValNode struct {
		row int
		col int
		val int //string bool interface
	}

	var SparseArr []ValNode

	SparseArr = append(SparseArr, ValNode{
		row: row,
		col: col,
		val: val,
	})

	for i, r := range Arr {
		for j, v := range r {
			if v != 0 {
				SparseArr = append(SparseArr, ValNode{
					row: i,
					col: j,
					val: v,
				})
			}
		}
	}
	for i, node := range SparseArr {
		fmt.Printf("%d: %d %d %d\n", i, node.row, node.col, node.val)
	}

	//将稀疏数组保存到文件
	filename := "sparsearr.data"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("写入失败")
	}
	for _, node := range SparseArr {
		str := fmt.Sprintf("%d %d %d\n", node.row, node.col, node.val)
		fmt.Println(str)
		_, err := io.WriteString(f, str)
		if err != nil {
			fmt.Printf("写入失败")
		}
	}

	//从文件中恢复二维数组
	f, err = os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件失败")
	}
	buf := bufio.NewScanner(f)
	mapRow := 0
	mapCol := 0
	//mapVal := 0
	for i := 0; ; i++ {
		if !buf.Scan() {
			break
		}
		if i == 0 {
			line := buf.Text()
			line = strings.TrimSpace(line)
			node := strings.Split(line, " ")
			mapRow, _ = strconv.Atoi(node[0])
			mapCol, _ = strconv.Atoi(node[1])
			fmt.Println(node)
			fmt.Println("row:", mapRow)
			fmt.Println("col:", mapCol)
			//mapVal, _ = strconv.Atoi(node[2])
		} else {
			line := buf.Text()
			node := strings.Split(line, " ")
			r, _ := strconv.Atoi(node[0])
			c, _ := strconv.Atoi(node[1])
			v, _ := strconv.Atoi(node[2])
			//Map[r][c] = v
			//TODO:恢复稀疏数组
			fmt.Printf("this: %d %d %d", r, c, v)
			fmt.Println()
		}
	}
	Map := make([][]int, mapRow)
	for i := range Arr {
		Map[i] = make([]int, mapCol)
	}
	for _, r := range Map {
		for _, v := range r {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
