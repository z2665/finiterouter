package tree

import (
	"net/http"
)

const MAXLENS = 26

type Node struct {
	Val     byte
	isfinal bool
	next    [MAXLENS]*Node
	Handle  http.HandlerFunc
}
type Tree struct {
	first *Node
}

func NewTree() *Tree {
	return &Tree{
		first: &Node{},
	}
}

//返回最后一位节点
func (t *Tree) Insert(s string) *Node {
	// if t.first == nil {
	// 	t.first = &Node{}
	// }
	index := t.first
	var p *Node
	for i := 0; i < len(s); i++ {
		c := toLowerAscii(s[i])
		p = index.next[c%MAXLENS]
		if p == nil {
			p = &Node{
				Val: c,
			}
			index.next[c%MAXLENS] = p
		}
		index = p
	}
	//此时的p为最后一位字符节点
	p.isfinal = true
	return p
}
func (t *Tree) Search(s string) *Node {
	index := t.first
	var p *Node
	for i := 0; i < len(s); i++ {
		c := toLowerAscii(s[i])
		p = index.next[c%MAXLENS]
		if p == nil {
			return nil
		}
		index = p
	}
	if p.isfinal != true {
		return nil
	}
	return p
}

func (t *Tree) First() *Node {
	return t.first
}

//假定ascii字符范围为字母集
func toLowerAscii(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		c += 'a' - 'A'
	}
	return c
}
