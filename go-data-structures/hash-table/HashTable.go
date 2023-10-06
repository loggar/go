package main

import "fmt"

// ArraySize is size of hash table array
const ArraySize = 5

// HashTable
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of hash table
type bucket struct {
	head *bucketNode
}

// bucketNode
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will add key-value to hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will return true if given key is stored in hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will delete node from hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will create a node with key and insert it in bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		node := &bucketNode{key: k}
		node.next = b.head
		b.head = node
	}
}

// search will return true if given key is in bucket
func (b *bucket) search(k string) bool {
	node := b.head
	for node != nil {
		if node.key == k {
			return true
		}
		node = node.next
	}
	return false
}

// delete will remove a node from bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	node := b.head
	for node.next != nil {
		if node.next.key == k {
			node.next = node.next.next
			return
		}
		node = node.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	fmt.Println(hash("RANDY"))

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	testBucket.insert("TOM")
	testBucket.insert("TOM")
	fmt.Printf("testBucket %v\n", testBucket)
	fmt.Printf("*testBucket.head %v\n", *testBucket.head)
	fmt.Printf("*testBucket.head.next %v\n", *testBucket.head.next)
	fmt.Printf("search TOM %v\n", testBucket.search("TOM"))
	fmt.Printf("search RANDY %v\n", testBucket.search("RANDY"))
	fmt.Printf("search JOHN %v\n", testBucket.search("JOHN"))

	testBucket.delete("TOM")
	fmt.Printf("search RANDY %v\n", testBucket.search("RANDY"))
	fmt.Printf("search TOM %v\n", testBucket.search("TOM"))

	hashTbl := Init()
	fmt.Println(hashTbl)

	list := []string{
		"ERIC",
		"CHARLY",
		"KENNY",
		"RANDY",
		"TOM",
	}

	for _, v := range list {
		hashTbl.Insert((v))
	}

	hashTbl.Delete("RANDY")

	fmt.Printf("Search CHARLY %v\n", hashTbl.Search("CHARLY"))
	fmt.Printf("Search RANDY %v\n", hashTbl.Search("RANDY"))
}
