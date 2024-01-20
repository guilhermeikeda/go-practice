package main

import "fmt"

type CacheNode struct {
	key, value int

	right *CacheNode
	left  *CacheNode
}

type LRUCache struct {
	cache map[int]*CacheNode

	capacity int

	start, end *CacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*CacheNode, capacity),
	}
}

func (c *LRUCache) addTop(node *CacheNode) {
	node.left = nil

	if c.start == nil {
		c.start = node
		c.end = node
	} else {
		node.right = c.start
		c.start.left = node

		c.start = node
	}

	fmt.Printf("%v \n", *c.start)
	fmt.Printf("%v \n", *c.end)
}

func (d *LRUCache) remove(node *CacheNode) {
	// Node is at the start
	if node.left == nil {
		d.start = node.right
	} else {
		node.left.right = node.right
	}

	if node.right == nil {
		d.end = node.left
	} else {
		node.right.left = node.left
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.remove(node)
		c.addTop(node)
		return node.value
	}

	return -1
}

func (c *LRUCache) Put(key, value int) {
	if node, ok := c.cache[key]; ok {
		node.value = value // update node value
		c.remove(node)     // increase priority
		c.addTop(node)
		return
	}

	node := CacheNode{
		key:   key,
		value: value,
	}

	if len(c.cache) >= c.capacity {
		// make room for new element, by removing the end node
		delete(c.cache, c.end.key)
		c.remove(c.end)
	}

	c.addTop(&node)
	c.cache[key] = c.start
}
