package fourth_edition

// 正则表达式
// 非确定有限状态自动机NFA
// 表示NFA：
// 括号：每当遇到一个右括号，我们最终都会将左括号从栈中弹出
// 闭包操作：*出现在单个字符之后或者右括号之后，
//	前者在该字符和*之间添加互相指向的两条€转换，后者在对应的左括号和*添加互相指向的两条€转换
// 或表达式：如（A|B）,添加两条€转换，一条从左括号指向B，另一条从|字符指向右括号
type NFA struct {
	re []byte // 匹配转换
	digraph *DirectedGraph // epsilon转换
	m int // 状态的数量
}

func NewNFA(regexp string) *NFA {
	ops := NewStack()
	re := []byte(regexp)
	m := len(re)
	g := new(DirectedGraph)
	g.New(m+1)

	for i:=0; i < m; i++ {
		lp := i
		if re[i] == '(' || re[i] == '|' {
			ops.push(i)
		}else if re[i] == ')' {
			or := ops.pop()
			if re[or] == '|' {
				lp = ops.pop() // 此时 lp为(，or为|，|指向)，(指向|的下一个字符
				g.AddEdge(lp, or+1)
				g.AddEdge(or, i)
			}else {
				lp = or // 此时or、lp为(
			}
		}
		if i < m-1 && re[i+1] == '*' { // 当前字符与*互相指向 或者 (与*互相指向
			g.AddEdge(lp, i+1)
			g.AddEdge(i+1, lp)
		}
		if re[i] == '(' || re[i] == '*' || re[i] == ')' { // (、*、)指向下一个字符
			g.AddEdge(i, i+1)
		}
	}
	return &NFA{
		re:      re,
		digraph: g,
		m:       m,
	}
}

func (n NFA) recognizes(txt string) bool {
	pc := make(map[int]struct{}) // 能够达到的元素集合
	dfs := new(DirectedDFS)
	dfs.Search(*n.digraph, 0)
	for v := 0; v <n.digraph.V(); v++ {
		if dfs.hasMarked(v) {
			pc[v] = struct{}{}
		}
	}

	for i := 0; i < len(txt); i++ {
		// 计算txt[i+1]可能到达的所有NFA状态
		match := make(map[int]struct{})
		for v := range pc {
			if v < n.m {
				if n.re[v] == txt[i] || n.re[v] == '.' {
					match[v+1] = struct{}{}
				}
			}
		}
		pc = make(map[int]struct{})
		dfs = new(DirectedDFS)
		dfs.dfsMap(*n.digraph, match)
		for v := 0; v < n.digraph.V(); v++ {
			if dfs.hasMarked(v) {
				pc[v] = struct{}{}
			}
		}
	}
	for v := range pc {
		if v == n.m {
			return true
		}
	}
	return false

}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) pop() int {
	if len(s.data) == 0 {
		return -1
	}
	result := s.data[0]
	s.data = s.data[1:]
	return result
}