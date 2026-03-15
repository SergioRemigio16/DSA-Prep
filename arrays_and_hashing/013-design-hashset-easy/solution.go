package main

import "container/list"

type MyHashSet struct {
	// size of the table
	m int
	// array slice that will hold keys
	table []*list.List
}

func Constructor() MyHashSet {
	m := 10
	myHashSet := MyHashSet{
		m:     m,
		table: make([]*list.List, m),
	}
	for i := 0; i < m; i++ {
		myHashSet.table[i] = list.New()
	}
	return myHashSet
}

func (this *MyHashSet) Add(key int) {
	index := this.hash(key)
	if !this.Contains(key) {
		this.table[index].PushFront(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	if !this.Contains(key) {
		return
	}
	index := this.hash(key)
	currNode := this.table[index].Front()
	// Iterate through doubly linked list and find key
	for {
		if currNode.Value == key {
			// remove key
			this.table[index].Remove(currNode)
			break
		}
		// Go to the nex node
		currNode = currNode.Next()
	}
}

func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)
	// Check if doubly linked list is empty
	if this.table[index].Len() == 0 {
		return false
	} else {
		currNode := this.table[index].Front()
		for {
			if currNode.Value == key {
				return true
			}
			if currNode.Next() == nil { // reached end of list
				break
			}
			currNode = currNode.Next()
		}
	}
	return false
}

// Private hash function
func (this *MyHashSet) hash(key int) int {
	return key % this.m
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
	myHashSet := Constructor()
	/* Test 1
	myHashSet.Add(1)
	myHashSet.Add(2)
	myHashSet.Contains(1)
	myHashSet.Contains(3)
	myHashSet.Add(2)
	myHashSet.Contains(2)
	myHashSet.Remove(2)
	println(myHashSet.Contains(2))
	*/
	// Test 2
	myHashSet.Add(10)
	myHashSet.Add(20)
	myHashSet.Add(30)
	myHashSet.Remove(10)
	myHashSet.Contains(10)
	myHashSet.Remove(10)
	myHashSet.Contains(10)
}
