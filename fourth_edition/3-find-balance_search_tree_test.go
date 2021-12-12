package fourth_edition

import (
	"math/rand"
	"testing"
)

func TestRedBlackBST(t *testing.T) {
	bst := new(RedBlackBST)
	for i := 0; i < 10; i++ {
		rd := Comparable(rand.Intn(100))
		bst.put(rd, rd)
	}
	bst.root.show()
}
