// https://github.com/google/btree
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type BTreeNode struct {
    t        int
    keys     []int
    children []*BTreeNode
    leaf     bool
}

func NewBTreeNode(t int, leaf bool) *BTreeNode {
    return &BTreeNode{
        t:        t,
        keys:     make([]int, 0),
        children: make([]*BTreeNode, 0),
        leaf:     leaf,
    }
}

func (n *BTreeNode) InsertNonFull(k int) {
    i := len(n.keys) - 1
    if n.leaf {
        n.keys = append(n.keys, 0)
        for i >= 0 && k < n.keys[i] {
            n.keys[i+1] = n.keys[i]
            i--
        }
        n.keys[i+1] = k
    } else {
        for i >= 0 && k < n.keys[i] {
            i--
        }
        i++
        if len(n.children[i].keys) == 2*n.t-1 {
            n.splitChild(i)
            if k > n.keys[i] {
                i++
            }
        }
        n.children[i].InsertNonFull(k)
    }
}

func (n *BTreeNode) splitChild(i int) {
    t := n.t
    y := n.children[i]
    z := NewBTreeNode(t, y.leaf)
    n.children = append(n.children, nil)
    copy(n.children[i+2:], n.children[i+1:])
    n.children[i+1] = z
    n.keys = append(n.keys, 0)
    copy(n.keys[i+1:], n.keys[i:])
    n.keys[i] = y.keys[t-1]
    z.keys = append(z.keys, y.keys[t:]...)
    y.keys = y.keys[:t-1]
    if !y.leaf {
        z.children = append(z.children, y.children[t:]...)
        y.children = y.children[:t]
    }
}

func (n *BTreeNode) print() {
    fmt.Printf("Keys: %v\n", n.keys)
    for _, child := range n.children {
        child.print()
    }
}

type BTree struct {
    root *BTreeNode
    t    int
}

func NewBTree(t int) *BTree {
    return &BTree{
        root: NewBTreeNode(t, true),
        t:    t,
    }
}

func (t *BTree) Insert(k int) {
    root := t.root
    if len(root.keys) == (2*t.t - 1) {
        temp := NewBTreeNode(t.t, false)
        temp.children = append(temp.children, root)
        t.root = temp
        temp.splitChild(0)
        temp.children[0].InsertNonFull(k)
    } else {
        root.InsertNonFull(k)
    }
}

func (t *BTree) Print() {
    t.root.print()
}

// BTree and BTreeNode are defined as before...

// Assume that each node is stored in a separate file, and the file name is the block ID.
func (n *BTreeNode) saveToFile(blockID string) error {
    file, err := os.Create(blockID)
    if err != nil {
        return err
    }
    defer file.Close()

    enc := gob.NewEncoder(file)
    if err := enc.Encode(n); err != nil {
        return err
    }

    return nil
}

func loadNodeFromFile(blockID string) (*BTreeNode, error) {
    file, err := os.Open(blockID)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var n BTreeNode
    dec := gob.NewDecoder(file)
    if err := dec.Decode(&n); err != nil {
        return nil, err
    }

    return &n, nil
}

func (n *BTreeNode) search(k int) (*BTreeNode, int, error) {
    i := 0
    for i < len(n.keys) && k > n.keys[i] {
        i++
    }

    if i < len(n.keys) && k == n.keys[i] {
        return n, i, nil
    } else if n.leaf {
        return nil, -1, fmt.Errorf("key not found")
    } else {
        child, err := loadNodeFromFile(n.children[i])
        if err != nil {
            return nil, -1, err
        }
        return child.search(k)
    }
}

func main() {
    // Assume that the BTree has been built and each node has been saved to a file...
    rootID := "root"
    root, err := loadNodeFromFile(rootID)
    if err != nil {
        fmt.Println(err)
        return
    }

    node, index, err := root.search(5)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Found key in node %v at index %d\n", node, index)
    }
}