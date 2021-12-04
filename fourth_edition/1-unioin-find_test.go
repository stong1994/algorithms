package fourth_edition

import "testing"

var ufMaxNum = 10
var ufData = [][2]int{
	{4, 3},
	{3, 8},
	{4, 8},
	{6, 5},
	{9, 4},
	{2, 1},
	{8, 9},
	{5, 0},
	{7, 2},
	{6, 1},
	{1, 0},
	{6, 7},
}

func TestUF_quickFind(t *testing.T) {
	uf := UF{
		id:    nil,
		count: 0,
		iuf:   &quickFind{},
	}
	ExecUF(&uf, ufMaxNum, ufData)
}

func TestUF_quickUnion(t *testing.T) {
	uf := UF{
		id:    nil,
		count: 0,
		iuf:   &quickUnion{},
	}
	ExecUF(&uf, ufMaxNum, ufData)
}

func TestUF_weightQuickUnion(t *testing.T) {
	uf := UF{
		id:    nil,
		count: 0,
		iuf:   &weightedQuickUnion{},
	}
	ExecUF(&uf, ufMaxNum, ufData)
}