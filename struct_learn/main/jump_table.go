package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data int
	node []*Node
}

type JumpTable struct {
	head   Node
	level  int
	length int
}

func New(level int) JumpTable {
	return JumpTable{head: createNode(0, level), level: level}
}

func createNode(data, level int) Node {
	nodes := make([]*Node, level)
	n := Node{data: data}
	for i, _ := range nodes {
		nodes[i] = nil
	}
	n.node = nodes
	return n
}

func (t *JumpTable) Put(data int) {
	head := t.head
	level := t.level
	updateNode := make([]*Node, t.level)
	current := head
	for i := level - 1; i >= 0; i-- {
		// 判断第一层
		for current.node[i] != nil && current.node[i].data < data {
			current = *current.node[i]
		}
		updateNode[i] = &current
	}
	// 插入成功
	newLevel := randomLevel(level)
	if newLevel < level {
		level = newLevel
	}
	//逐步查询
	n := Node{data: data, node: make([]*Node, newLevel)}
	// 稀疏程度 通过"随机产生层数来进行，随机稀疏度
	for i := level - 1; i >= 0; i-- {
		// 是否维持稀疏度？
		if len(updateNode[i].node)>i {
			n.node[i] = updateNode[i].node[i]
			updateNode[i].node[i] = &n
		}
	}
	fmt.Println(newLevel,updateNode,n)
	t.length += 1
}

func randomLevel(level int) int {
	return 1 + rand.Intn(level)
}

// 1   4   8
// 1 2 4 6 8
func (t *JumpTable) Pop(data int) {
	for i := t.level - 1; i > 0; i-- {

	}
}

func (t *JumpTable) For() {
	current := &t.head
	fmt.Println(current)
	for current.node[0] != nil {
		current = current.node[0]
		fmt.Print(current.data," ")
	}
	fmt.Println(" ")
	current = &t.head
	index := 1
	fmt.Println(current)
	for {
		if len(current.node)>1 && current.node[1] !=nil {
			index=1
			current = current.node[index]
			fmt.Print("|")
		}else{
			index=0
			current = current.node[index]
		}
		if current==nil {
			break
		}
		fmt.Print(current.data," ")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	n := New(3)
	for i := 1; i < 16; i++ {
		n.Put(i)
	}
	n.For()
}
