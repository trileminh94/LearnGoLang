package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	root := &node{make(map[byte]*node), false}
	root.data['a'] = &node{make(map[byte]*node), false}
	root.isFinalNode = true

	root.data['a'].data['b'] = &node{make(map[byte]*node), true}

	type args struct {
		n *node
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test search ab", args{root, "ab"}, true},
		{"test search ac", args{root, "ac"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_addchar(t *testing.T) {
	type fields struct {
		data        map[byte]*node
		isFinalNode bool
	}
	type args struct {
		a byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"test add a", fields{make(map[byte]*node), false}, args{'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				data:        tt.fields.data,
				isFinalNode: tt.fields.isFinalNode,
			}
			n.addchar(tt.args.a)

			_, ok := n.data['a']
			assert.True(t, ok)
			assert.True(t, n.isFinalNode)
		})
	}
}

func Test_node_add(t *testing.T) {
	type fields struct {
		data        map[byte]*node
		isFinalNode bool
	}
	type args struct {
		a string
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"add abc find abc", fields{make(map[byte]*node), false}, args{"abc", "abc"}, true},
		{"add abc find ab", fields{make(map[byte]*node), false}, args{"abc", "ab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				data:        tt.fields.data,
				isFinalNode: tt.fields.isFinalNode,
			}
			n.add(tt.args.a)
			if got := search(n, tt.args.s); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
