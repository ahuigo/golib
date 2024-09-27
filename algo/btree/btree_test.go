package main

import (
	"fmt"
	"testing"

	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"
)
//  btree　内存结构(不带持久化)
func TestBtreeCrud(t *testing.T) {
	tree := btree.NewWith(3, utils.IntComparator)

	// Insert
	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Put(3, "c")

	// Find
	val, found := tree.Get(2)
	if found {
		fmt.Println(val)
	}

	// Delete
	tree.Remove(2)

	// Check if deleted
	_, found = tree.Get(2)
	if !found {
		fmt.Println("Value 2 not found")
	}
}

