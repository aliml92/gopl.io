package ex73

import (
	"strconv"
	"strings"
)


type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}


// String return in-order traversal representation of tree with arrows, e.g. 0->1->2

func (t *tree) String() string {
	var dfs func(root *tree)
	var s strings.Builder
	dfs = func(root *tree) {
		if root == nil {
			return 
		}
		dfs(root.left)
		s.WriteString(strconv.Itoa(root.value))
		if s.Len() > 0 && s.String()[s.Len()-1:] != ">" {
			s.WriteString("->")
		}
		dfs(root.right)
	}
	dfs(t)
	if s.Len() > 0 {
		return s.String()[:s.Len()-2]
	}
	return s.String()
}