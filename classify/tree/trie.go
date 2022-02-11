package tree

// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
//请你实现 Trie 类：
//Trie() 初始化前缀树对象。
//void insert(String word) 向前缀树中插入字符串 word 。
//boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
//boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
// 1 <= word.length, prefix.length <= 2000
// word 和 prefix 仅由小写英文字母组成
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		if node.children[v] == nil {
			node.children[v] = new(Trie)
		}
		node = node.children[v]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	L := len(word)
	for i := 0; i < L-1; i++ {
		v := word[i] - 'a'
		if node.children[v] == nil {
			return false
		}
		node = node.children[v]
	}
	last := node.children[word[L-1]-'a']
	if last == nil {
		return false
	}
	return last.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		v := prefix[i] - 'a'
		if node.children[v] == nil {
			return false
		}
		node = node.children[v]
	}
	return true
}

// 实现一个 Trie，用来求前缀和
// 设计一个 map ，满足以下几点:
//	字符串表示键，整数表示值
//	返回具有前缀等于给定字符串的键的值的总和
//	实现一个 MapSum 类：
//	MapSum() 初始化 MapSum 对象
//	void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。
//	如果键 key 已经存在，那么原来的键值对key-value将被替代成新的键值对。
//	int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
// 提示：
//	1 <= key.length, prefix.length <= 50
//	key 和 prefix 仅由小写英文字母组成
//	1 <= val <= 1000
//	最多调用 50 次 insert 和 sum
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/map-sum-pairs
type TrieNode struct {
	children [26]*TrieNode
	val      int
}
type MapSum struct {
	node *TrieNode
	cnt  map[string]int
}

func ConstructorMapSum() MapSum {
	return MapSum{node: new(TrieNode), cnt: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val - this.cnt[key]
	this.cnt[key] = val
	node := this.node
	for _, v := range key {
		n := v - 'a'
		if node.children[n] == nil {
			node.children[n] = &TrieNode{val: delta}
		} else {
			node.children[n].val += delta
		}
		node = node.children[n]
	}
}

func (this *MapSum) Sum(prefix string) int {
	node := this.node
	for _, v := range prefix {
		n := v - 'a'
		if node.children[n] == nil {
			return 0
		}
		node = node.children[n]
	}
	return node.val
}
