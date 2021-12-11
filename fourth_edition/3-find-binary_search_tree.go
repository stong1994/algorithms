package fourth_edition

// 在查找领域:
// hash的查找与插入效率都很高，但无法进行顺序相关操作。
// 有序数组查找很快，但插入很慢
// 链表插入很快，但查找很慢
// 二叉查找树则综合了二分查找的效率与链表的灵活性，但是没有性能上界的保证
// 平衡二叉查找树则能够有最优的查找和插入效率，同时能够进行有序的相关操作。
// 不同于堆，树的父节点大于左子节点，小于右子节点

type Node struct {
	key         Comparable
	value       Comparable
	left, right *Node // 两个子节点
	n           int   // 以该节点为根节点的节点总数
}

func NewNode(key, value Comparable, left, right *Node, n int) *Node {
	return &Node{
		key:   key,
		value: value,
		left:  left,
		right: right,
		n:     n,
	}
}

// 二叉查找树
type BinarySearchTree struct {
	root *Node
}

func (b BinarySearchTree) get(key Comparable) Comparable {
	return b.getNode(b.root, key)
}

func (b BinarySearchTree) getNode(node *Node, key Comparable) Comparable {
	if node == nil {
		return comparableNull
	}
	cmp := node.key.CompareTo(key)
	if cmp == 0 {
		return node.value
	}
	if cmp > 0 {
		return b.getNode(node.right, key)
	}
	return b.getNode(node.left, key)
}

func (b *BinarySearchTree) put(key, value Comparable) {
	b.root = b.updateOrCreate(b.root, key, value)
}

// 查找到则更新，否则新增
func (b *BinarySearchTree) updateOrCreate(node *Node, key, value Comparable) *Node {
	if node == nil {
		return NewNode(key, value, nil, nil, 1)
	}
	cmp := node.key.CompareTo(key)
	if cmp == 0 {
		node.value = value
		return node
	}
	if cmp > 0 {
		node.right = b.updateOrCreate(node.right, key, value)
	} else {
		node.left = b.updateOrCreate(node.left, key, value)
	}
	node.n = b.size(node.left) + b.size(node.right) + 1
	return node
}

func (b BinarySearchTree) min(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return b.min(node.left)
}

// 小于等于key
func (b BinarySearchTree) floor(key Comparable) Comparable {
	n := b.floorNode(b.root, key)
	if n == nil {
		return comparableNull
	}
	return n.value
}

func (b BinarySearchTree) floorNode(node *Node, key Comparable) *Node {
	if node == nil {
		return nil
	}
	cmp := node.key.CompareTo(key)
	if cmp == 0 {
		return node
	}
	if cmp > 0 {
		return b.floorNode(node.left, key)
	}
	r := b.floorNode(node.right, key)
	if r == nil {
		return node
	}
	return r
}

// 找到排名为k的节点
func (b BinarySearchTree) selectNode(k int) *Node {
	return b.selectWithNode(b.root, k)
}

func (b BinarySearchTree) selectWithNode(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	if node.n == k {
		return node
	}
	if node.n > k {
		return b.selectWithNode(node.left, k)
	}
	return b.selectWithNode(node.right, k-node.n)
}

func (b *BinarySearchTree) deleteMin() {
	b.root = b.deleteMinNode(b.root)
}

func (b *BinarySearchTree) deleteMinNode(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = b.deleteMinNode(node.left)
	node.n = b.size(node.left) + b.size(node.right) + 1
	return node
}

func (b *BinarySearchTree) size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.n
}

func (b *BinarySearchTree) delete(key Comparable) {
	b.root = b.deleteNode(b.root, key)
}

func (b *BinarySearchTree) deleteNode(node *Node, key Comparable) *Node {
	if node == nil {
		return nil
	}
	cmp := node.key.CompareTo(key)
	if cmp < 0 {
		node.right = b.deleteNode(node.right, key)
		return node
	}
	if cmp > 0 {
		node.left = b.deleteNode(node.left, key)
		return node
	}
	if node.right == nil {
		return node.left
	}
	if node.left == nil {
		return node.right
	}
	r := b.deleteMinNode(node.right)
	r.left = node.left
	r.right = node.right
	r.n = b.size(node.left) + b.size(node.right) + 1
	return r
}
