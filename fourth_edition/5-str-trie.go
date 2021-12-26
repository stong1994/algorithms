package fourth_edition

// 单词查找树——trie树：根节点为空，二级节点为单词的首字母节点，只在单词末尾所在的节点记录单词。
type TrieST struct {
	R    int       // 基数
	root *TrieNode // 单词查找树的根节点
}

func NewTrieST() *TrieST {
	return &TrieST{
		R:    256,
		root: nil,
	}
}

type TrieNode struct {
	val  int
	next []*TrieNode // 每个节点都是一个R长度的数组
}

func (s TrieST) NewTrieNode() *TrieNode {
	t := make([]*TrieNode, s.R)
	return &TrieNode{
		val:  0,
		next: t,
	}
}

func (s TrieST) get(key string) int {
	x := s.getNode(s.root, key, 0)
	if x == nil {
		return 0
	}
	return x.val
}

func (s TrieST) getNode(x *TrieNode, key string, d int) *TrieNode {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := charAt(key, d)
	return s.getNode(x.next[c], key, d+1)
}

func (s *TrieST) put(key string, val int) {
	s.root = s.putNode(s.root, key, val, 0)
}

func (s *TrieST) putNode(x *TrieNode, key string, val, d int) *TrieNode {
	if x == nil {
		x = s.NewTrieNode()
	}
	if d == len(key) {
		x.val = val
		return x
	}
	c := charAt(key, d)
	x.next[c] = s.putNode(x.next[c], key, val, d+1)
	return x
}

func (s TrieST) collect(x *TrieNode, pre string, q *queueString) {
	if x == nil {
		return
	}
	if x.val != 0 {
		q.enqueue(pre)
	}
	for c := 0; c < s.R; c++ {
		p := append([]byte(pre), byte(c))
		s.collect(x.next[c], string(p), q)
	}
}

func (s TrieST) keyWithPrefix(pre string) []string {
	q := newQueueString()
	s.collect(s.getNode(s.root, pre, 0), pre, q)
	return q.list
}

// 匹配模式
func (s TrieST) keysThatMatch(pat string) []string {
	q := newQueueString()
	s.collectThatMatch(s.root, "", pat, q)
	return q.list
}

func (s TrieST) collectThatMatch(x *TrieNode, pre, pat string, q *queueString) {
	if x == nil {
		return
	}
	d := len(pre)
	if d == len(pat) && x.val != 0 {
		q.enqueue(pre)
		return
	}
	if d == len(pat) {
		return
	}
	next := charAt(pat, d)
	for c := 0; c < s.R; c++ {
		if next == '.' || next == c {
			p := append([]byte(pre), byte(c))
			s.collectThatMatch(x.next[c], string(p), pat, q)
		}
	}
}

// 最长前缀
func (s TrieST) longestPrefixOf(str string) string {
	length := s.search(s.root, str, 0, 0)
	return str[0:length]
}

func (s TrieST) search(x *TrieNode, str string, d int, length int) int {
	if x == nil {
		return length
	}
	if x.val != 0 {
		length = d
	}
	if d == len(str) {
		return length
	}
	c := charAt(str, d)
	return s.search(x.next[c], str, d+1, length)
}
