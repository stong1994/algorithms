package fourth_edition

import "fmt"

// 平衡查找树
// 2-3查找树：含有两种节点
// 		2-节点：含有一个键和两条链接
// 		3-节点：含有两个键和三条链接
// 全局有序性和平衡性：任何空链接到根节点的路径长度都是相等的
// 插入规则：
//		向2-节点中插入新键，则直接变成3-节点
//		向3-节点中插入新键，则该节点临时变为4-节点，找到中键，并插入到父节点中。以此类推，直到找到一个2-节点。
//			如果从该节点到根节点都是3-节点，那么根节点变为临时的4-节点，然后分解为3-节点
//

// 实现：红黑树
// 红链接将两个2-节点连接起来构成一个3-节点；黑链是2-3树中的普通链接
// => 红黑树特征：
// 		- 红链接均为左链接
// 		- 没有任何一个节点同时和两条红链接相连
// 		- 完美黑色平衡，任意空链接到根节点的路径上的黑链接数量相同
// 如果将一个红黑树中的红链接画平，那么所有空链接到根节点的距离都是相同的
// 如果将一颗红黑树中的红链接相连的节点合并，那么得到的就是一颗2-3树
// 红黑树既是二叉查找树，也是2-3树 =》同时拥有二叉查找树中简洁高效的查找方法和2-3树种高效的平衡插入算法

type Color bool

func (c Color) String() string {
	if c == Red {
		return "red"
	}
	return "black"
}

const (
	Red   Color = true
	Black Color = false
)

type RBNode struct {
	key         Comparable
	value       Comparable
	left, right *RBNode
	n           int   // 这颗子树的节点总数
	color       Color // 其父节点指向它的链接的颜色，true=红色 false=黑色
}

func NewRBNode(key, value Comparable, n int, color Color) *RBNode {
	return &RBNode{
		key:   key,
		value: value,
		n:     n,
		color: color,
	}
}

func (h *RBNode) isRed() bool {
	if h == nil { // 空链为黑色
		return false
	}
	return h.color == Red
}

func RBNodeSize(h *RBNode) int {
	if h == nil {
		return 0
	}
	return h.n
}

// 旋转
// 某些操作可能会出现红色右链接或者两条连续的红链接，这时需要旋转
// 旋转操作可以保持红黑树的有序性和完美平衡性
// 左旋转：将红色右链接转化为红色左链接——将两个键中较小者作为根节点变为较大者作为根节点
func (h *RBNode) rotateLeft() *RBNode {
	newRoot := h.right
	h.right = newRoot.left
	newRoot.left = h
	newRoot.color = h.color // 新的根节点继承旧根节点的颜色
	h.color = Red
	newRoot.n = h.n
	h.n = 1 + RBNodeSize(h.left) + RBNodeSize(h.right)
	return newRoot
}

func (h *RBNode) rotateRight() *RBNode {
	newRoot := h.left
	h.left = newRoot.right
	newRoot.right = h
	newRoot.color = h.color // 新的根节点继承旧根节点的颜色
	h.color = Red
	newRoot.n = h.n
	h.n = 1 + RBNodeSize(h.left) + RBNodeSize(h.right)
	return newRoot
}

// 向单个2-节点插入新键
// 		如果新键小于老键，那么新增一个红色节点即可。这两个2-节点“组成”一个3-节点
// 		如果新键大于老键，那么新增的红色节点将产生一条红色的右链接，需要通过rotateLeft()来旋转
// 向树底部的2-节点插入新键（同上）
// 		如果指向新节点的父节点是左链接，那么父节点直接变为3-节点
// 		如果指向新节点的父节点是右链接，那么父节点将变为错误的3-节点，，需要一次左旋转修复
// 向3-节点（一棵双键树）插入新键
// 		如果新键大于原树中的两个键：将它连接到3-节点的右链接（此时树是平衡的，根节点是中间大小的键），
// 		如果新键小于原树中的两个键：将它连接到最左边的空链接，这样产生了两条连续的红链接，需要将上层的红链接右旋
// 		如果新键位于原树中的两个键之间：将它连接到”左节点“的右链接，然后左旋”左节点“，然后再左旋”根节点“
// 		这三种情况最后倒要将两条链接变为黑色（此时这棵树变为了2-3树）然后将“根节点”的颜色置为红色
// 		这时候“根节点”与其父节点组成了3-节点或者4-节点
// 			如果是3-节点，那么处理方式同“向单个2-节点插入新键”，
//			如果是4-节点，那么处理方式同“向3-节点（一棵双键树）插入新键”

// 颜色变换规律：
// 		如果右子节点是红色而左子节点是黑色，则进行左旋转
// 		如果左子节点是红色而该节点也是红色，则进行右旋转
// 		如果左右子节点都是红色，则进行颜色转换

// 将父节点的颜色由黑变红，将子节点变为黑色
func (h *RBNode) flipColors() {
	h.color = Red
	h.left.color = Black
	h.right.color = Black
}

func (h *RBNode) show() {
	if h == nil {
		return
	}
	root := []*RBNode{h}
	for root != nil {
		for _, v := range root {
			if v == nil {
				fmt.Print(" null | ")
				continue
			}
			fmt.Print("k:", v.key, " n:", v.n, " color:", v.color, " | ")
		}
		fmt.Println()
		root = h.nextLayerNode(root)
	}
}

func (h RBNode) nextLayerNode(ns []*RBNode) []*RBNode {
	var result []*RBNode
	for _, v := range ns {
		if v != nil {
			result = append(result, v.left, v.right)
		}
	}
	return result
}

type RedBlackBST struct {
	root *RBNode
}

func (r *RedBlackBST) put(key, value Comparable) {
	r.root = r.putNode(r.root, key, value)
	r.root.color = Black
}

func (r *RedBlackBST) putNode(h *RBNode, key, value Comparable) *RBNode {
	if h == nil {
		return NewRBNode(key, value, 1, Red) // 和父节点用红链接相连
	}
	cmp := key.CompareTo(h.key)
	if cmp > 0 {
		h.right = r.putNode(h.right, key, value)
	} else if cmp < 0 {
		h.left = r.putNode(h.left, key, value)
	} else {
		h.value = value
	}
	if h.right.isRed() && !h.left.isRed() {
		h = h.rotateLeft()
	}
	if h.left.isRed() && h.left.left.isRed() {
		h = h.rotateRight()
	}
	if h.left.isRed() && h.right.isRed() {
		h.flipColors()
	}
	h.n = RBNodeSize(h.left) + RBNodeSize(h.right) + 1
	return h
}
