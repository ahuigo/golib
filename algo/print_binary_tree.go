/**
 * @author ahuigo
 * @test $ go run print_binary_tree.go
 */
package main

import "fmt"
import "math/rand"

func main() {
    // Print Binary Tree!!
	PrintBT(New(1), 5)
}

type T struct{
    b int  `json:"someStringFieldsName"`
}

// Walk walks the tree t sending all values
func PrintBT(t interface{}, max_pad uint){
    //print("1234567890123456789012345678901234567890123456789012345678901234567890\n")
    //print("        10        20        30        40        50        60        70\n")

	var a = &[]Que{{t,1 << max_pad, 1 << max_pad}}
	count:=0
	for len(*a)>0 {
		count++
		var b = []Que{}
        sline := ""
        line := ""
		for _,root := range *a{
            left_pad,right_pad := root.left_pad, root.right_pad
            node := root.node.(*Tree)
            Left := node.Left
            Right := node.Right

            if Left != nil {
                b = append(b, Que{Left, left_pad-right_pad/2, right_pad/2})
                sline = StrPad(sline, 0, left_pad - right_pad/4) + "/"
            }

            if Right != nil {
                b = append(b, Que{Right, left_pad+right_pad/2, right_pad/2})
                sline = StrPad(sline, 0, left_pad + right_pad/4) + "\\"
            }

            line = StrPad(line, 0, left_pad) + fmt.Sprintf("%v", node.Value) 
		}
		print(line, "\n")
		print(sline, "\n")
		a = &b
	}
}

type Que struct {
	node  interface{}
	left_pad int
    right_pad int
}
func StrPad(s interface{}, left int, right int) string{
    var format string
    var str string
    str = fmt.Sprintf("%v", s)
    if left>0{
        format = fmt.Sprintf("%%%dv", left)
        str = fmt.Sprintf(format, str)
    }
    if right>0{
        format = fmt.Sprintf("%%-%vv", right)
        str = fmt.Sprintf(format, str)
    }
    return str
}

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(15) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

/*
func (t []*T) String() string {
    var s string
    for _,v := range *t{
        s = s + v.String()
    }
    return s
}
*/
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
