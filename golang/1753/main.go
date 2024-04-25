package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Node struct {
	adj map[int]int
}

type Item struct {
	number   int
	distance int
	idx      int
}

type PriorityQueue []*Item

var pq PriorityQueue

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	l := len(*pq)
	newItem := x.(*Item)
	newItem.idx = l
	*pq = append(*pq, newItem)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq

	l := len(old)
	item := old[l-1]
	old[l-1] = nil
	item.idx = -1
	*pq = old[0 : l-1]
	return item
}

func getShort(isVisited []bool, distance []int, nodeList []Node, num int) {
	heap.Init(&pq)

	start := &Item{number: num, distance: 0, idx: 0}
	heap.Push(&pq, start)

	for pq.Len() > 0 {
		top := heap.Pop(&pq)
		number := top.(*Item).number
		len := top.(*Item).distance

		if isVisited[number] == true {
			continue
		}

		isVisited[number] = true
		distance[number] = len

		for to, w := range nodeList[number].adj {
			if isVisited[to] == true {
				continue
			}

			newItem := &Item{to, len + w, 0}
			heap.Push(&pq, newItem)
		}
	}
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	var a, b, s int

	defer writer.Flush()

	fmt.Fscan(reader, &a, &b)
	fmt.Fscan(reader, &s)

	distance := make([]int, a+1)
	nodeList := make([]Node, a+1)
	isVisited := make([]bool, a+1)

	for i := 1; i <= a; i++ {
		nodeList[i].adj = make(map[int]int)
	}

	for i := 1; i <= b; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		if nodeList[u].adj[v] == 0 {
			nodeList[u].adj[v] = w
		} else if w < nodeList[u].adj[v] {
			nodeList[u].adj[v] = w
		}
	}

	getShort(isVisited, distance, nodeList, s)

	for i := 1; i <= a; i++ {
		d := distance[i]
		if i == s {
			fmt.Fprintln(writer, 0)
		} else {
			if d == 0 {
				fmt.Fprintln(writer, "INF")
			} else {
				fmt.Fprintln(writer, d)
			}
		}
	}

}
