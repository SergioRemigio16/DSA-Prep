package main

import "container/list"

// MyHashMap() initializes the object with an empty map.
type MyHashMap struct {
	// size of the table
	m int
	// array slice that will hold keys
	table []*list.List
}

type cell struct {
	key   int
	value int
}

func Constructor() MyHashMap {
	m := 10
	myHashSet := MyHashMap{
		m:     m,
		table: make([]*list.List, m),
	}
	for i := 0; i < m; i++ {
		myHashSet.table[i] = list.New()
	}
	return myHashSet
}

// void put(int key, int value) inserts a (key, value) pair into the HashMap.
// If the key already exists in the map, update the corresponding value.
func (this *MyHashMap) Put(key int, value int) {
	index := this.hash(key)
	currNode := this.table[index].Front()
	if currNode == nil {
		c := cell{}
		c.key = key
		c.value = value
		this.table[index].PushFront(c)
		return
	}
	for {
		if currNode.Value.(cell).key == key {
			c := currNode.Value.(cell)
			c.value = value
			currNode.Value = c
			return
		}
		if currNode.Next() == nil { // reached end of list
			break
		}
		currNode = currNode.Next()
	}
	c := cell{}
	c.key = key
	c.value = value
	this.table[index].PushFront(c)
}

// int get(int key) returns the value to which the specified key is mapped,
// or -1 if this map contains no mapping for the key.
func (this *MyHashMap) Get(key int) int {
	index := this.hash(key)
	if this.table[index].Len() == 0 {
		return -1
	} else {
		currNode := this.table[index].Front()
		for {
			if currNode.Value.(cell).key == key {
				return currNode.Value.(cell).value
			}
			if currNode.Next() == nil { // reached end of list
				break
			}
			currNode = currNode.Next()
		}
	}
	return -1
}

// void remove(key) removes the key and its corresponding value if the
// map contains the mapping for the key.
func (this *MyHashMap) Remove(key int) {
	index := this.hash(key)
	if this.table[index].Len() == 0 {
		return
	} else {
		currNode := this.table[index].Front()
		for {
			if currNode.Value.(cell).key == key {
				this.table[index].Remove(currNode)
			}
			if currNode.Next() == nil { // reached end of list
				break
			}
			currNode = currNode.Next()
		}
	}
}

// Private hash function
func (this *MyHashMap) hash(key int) int {
	return key % this.m
}

func main() {
	myHashMap := Constructor()
	myHashMap.Put(1, 1)
	myHashMap.Put(2, 2)
	myHashMap.Get(1)
	myHashMap.Get(3)
	myHashMap.Put(2, 1)
	myHashMap.Get(2)
	myHashMap.Remove(2)
	myHashMap.Get(2)
}
