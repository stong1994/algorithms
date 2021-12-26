package fourth_edition

// 三向单词查找树
// 每个节点都含有一个字符、三条链接和一个值，三条链接分别对应着当前字母小于、等于和大于节点字母的所有键
// trie树每个字母都要使用一个R长度数组，造成过度的空间消耗。
type TST struct {
	root *TSTNode
}

type TSTNode struct {
	c                int      // 字符
	left, mid, right *TSTNode // 左中右三个单词查找树
	val              int      // 存储值
}

func (t TST) get(x *TSTNode, key string, d int) *TSTNode {
	if x == nil {
		return nil
	}
	c := charAt(key, d)
	if c < x.c {
		return t.get(x.left, key, d)
	}
	if c > x.c {
		return t.get(x.right, key, d)
	}
	if d < len(key)-1 {
		return t.get(x.mid, key, d+1)
	}
	return x
}

func (t *TST) put(key string, val int) {
	t.root = t.putNode(t.root, key, val, 0)
}

func (t *TST) putNode(x *TSTNode, key string, val, d int) *TSTNode {
	c := charAt(key, d)
	if x == nil {
		x = new(TSTNode)
		x.c = c
	}
	if c < x.c {
		x.left = t.putNode(x.left, key, val, d)
	} else if c > x.c {
		x.right = t.putNode(x.right, key, val, d)
	} else if d < len(key)-1 {
		x.mid = t.putNode(x.mid, key, val, d+1)
	} else {
		x.val = val
	}
	return x
}
