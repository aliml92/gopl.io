package ex73

import (
	"math/rand"
	"sort"
	"testing"
)


func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}


func TestValues(t *testing.T) {
	tests := []struct{
		name string
		input *tree
		want string
	}{
		{name: "empty", input: nil, want: ""},
		{name: "duplicate", input: constuctTree(5, []int{0,1,8,9}), want: "0->1->5->8->9"},
		{name: "unique", input: constuctTree(0, []int{6,-4,6,15,11}), want: "-4->0->6->6->11->15"},
	}
	for _, tc := range tests {
		got := tc.input.String()
		if got != tc.want {
		t.Fatalf("got %v, want %v", got, tc.want)
	}
	}
}

func constuctTree(root int, children []int) *tree {
	tree := &tree{root, nil, nil}
	for _, child := range children {
		add(tree, child)
	}
	return tree 
}