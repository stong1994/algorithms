package str

// 回环字符串

/*
推导过程
*/

func findPalindrome(s string, l, r int) string {
	for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
		if s[l] != s[r] {
			break
		}
	}
	return s[l+1 : r]
}

func findLongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var result string
	for i := 0; i < len(s)-1; i++ {
		s1 := findPalindrome(s, i, i)
		s2 := findPalindrome(s, i, i+1)
		if len(s1) < len(s2) {
			s1 = s2
		}
		if len(s1) > len(result) {
			result = s1
		}
	}
	return result
}

// 计算一组字符集合可以组成的回文字符串的最大长度
// 给定一个包含大写字母和小写字母的字符串s，返回通过这些字母构造成的 最长的回文串。
// 在构造过程中，请注意 区分大小写 。比如"Aa"不能当做一个回文字符串。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindrome
func longestPalindrome(s string) int {
	// 最长回文需要保证左右对称，结果为偶数个数的元素的个数之和加上最多一个奇数个数的元素的数量
	// 需要注意奇数个数的元素也可以舍弃一个从而变为偶数个数，因此对于奇数个数的元素，也要加上其(cnt-1)的数量
	// 特别注意如果有元素的个数为1,3,5，那么需要加上4(5-1)、2(3-1)、1的数量，因为我们允许将1个数量的元素放到中间，但是这需要记录是否存在数量为1的元素
	// 为了方便代码，第一个个数为奇数的元素取其cnt值，后续的奇数数量的元素取其cnt-1值即可
	hash := make(map[int32]int)
	for _, v := range s {
		hash[v]++
	}
	var (
		result int
		hasOdd bool
	)
	for _, v := range hash {
		if v%2 == 0 {
			result += v
		} else {
			if !hasOdd {
				result += v
				hasOdd = true
			} else {
				result += v - 1
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 回文子字符串个数
// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
// 回文字符串 是正着读和倒过来读一样的字符串。
// 子字符串 是字符串中的由连续字符组成的一个序列。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindromic-substrings
func countSubstrings(s string) int {

	/* 方法一：中心扩展，对于每个元素来说，都可能是某个回文子串的中心，因此以此为中点，向两边进行扩展，并计数
	 回文的中心可能是偶数也可能是奇数，因此对于第i个元素，其可能是奇数回文的中心，也可能是偶数回文的中心的左元素，因此需要考虑两种情况
	var cnt int
	L := len(s)
	var match func(l, r int)
	match = func(l, r int) {
		if l < 0 || r >= L {
			return
		}
		if s[l] == s[r] {
			cnt++
			match(l-1, r+1)
		}
	}

	for i := 0; i < L-1; i++ {
		match(i, i)
		match(i, i+1)
	}
	cnt++ // 最后一个字符为回文中心的情况
	return cnt
	*/

	/*
		// 上述方案的优化版本：观察可发现对于含有L个元素的字符串，最多有2L-1个回文中心，分别是（0,0）（0，1），（1,1），（1,2）。。（L-2,L-2）,(L-2,L-1),(L-1,L-1)
		// 即对于第i个回文中心，其左中心索引为i/2,右中心索引为i/2+i%2
		var cnt int
		L := len(s)
		for i := 0; i < 2*L-1; i++ {
			left, right := i/2, i/2+i%2
			for ;left>= 0 && right < L && s[left] == s[right]; left, right = left-1, right+1 {
				cnt++
			}
		}
		return cnt
	*/

	newStr := preProcess(s)
	var (
		N      = len(newStr)
		p      = make([]int, N) // newStr对应的数组，元素值为在newStr中以当前元素为中心的回文字符串的最大半径。那么p[i]-1就是原字符串s以i为中心的最大回文长度
		rMax   = 0
		mid    = 0
		result = 0
	)
	for i := 1; i < N-1; i++ {
		if rMax > i {
			p[i] = min(rMax-i, p[mid*2-i])
		} else {
			p[i] = 1
		}
		for newStr[i+p[i]] == newStr[i-p[i]] {
			p[i]++
		}
		if i+p[i] > rMax {
			mid = i
			rMax = i + p[i]
		}
		result += p[i] / 2 // p[i]-1为第i个元素作为中心对应的回环长度，那么数量即为其除以二的向上取整值，即p[i]/2
	}
	return result
}

// 判断一个整数是否是回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
// 进阶：你能不将整数转为字符串来解决这个问题吗？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindrome-number
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 方法1，转成字符串，然后判断
	/*
		// 方法2，转成数组，然后判断
		var list []int
		for x > 0 {
			list = append(list, x%10)
			x /= 10
		}
		for left, right := 0, len(list)-1; left < right; left, right = left+1, right-1 {
			if list[left] != list[right] {
				return false
			}
		}
		return true
	*/
	// 方法3：方法2需要对每个数字进行除10，可以翻转一半的数字，再直接比较另一半来判断
	// 如果数字长度为偶数，那么当翻转一半后，翻转数字应等于原值；如果长度为奇数，那么翻转后，翻转数/10应等于原值
	reverseNum := 0
	for reverseNum < x {
		reverseNum = reverseNum*10 + x%10
		x /= 10
	}
	return reverseNum == x || reverseNum/10 == x
}

// 统计二进制字符串中连续 1 和连续 0 数量相同的子字符串个数
// 给定一个字符串s，统计并返回具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是成组连续的。
// 重复出现（不同位置）的子串也要统计它们出现的次数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/count-binary-substrings
func countBinarySubstrings(s string) int {
	/*方法一：通过中心向外扩展
	// 具有回文字符串的性质：从左读和从右读相同。在这基础上增加了条件：所有的0和所有的1都是连续的，也就是一半是0，另一半是1
	// 找到0和1相连的子串01，然后向外扩展，找到以01为中心的回环数量
	cnt := 0
	L := len(s)
	for i := 0; i < L-1; i++ {
		if s[i] == '0' && s[i+1] == '1' {
			cnt++
			for l, r := i-1, i+2; l >= 0 && r < L; l, r = l-1, r+1 {
				if s[l] == '0' && s[r] == '1' {
					cnt++
				}else {
					break
				}
			}
		}
		if s[i] == '1' && s[i+1] == '0' {
			cnt++
			for l, r := i-1, i+2; l >= 0 && r < L; l, r = l-1, r+1 {
				if s[l] == '1' && s[r] == '0' {
					cnt++
				}else {
					break
				}
			}
		}
	}
	return cnt
	*/
	/*
		// 方法二：将字符串转成相同字符数量的数组，相邻两个01能够组成的回环数取决于两者最小值
		var list []int
		lastNum := s[0]
		lastCount := 1
		for i := 1; i < len(s); i++ {
			if s[i] != lastNum {
				list = append(list, lastCount)
				lastCount = 1
				lastNum = s[i]
			}else {
				lastCount++
			}
		}
		list = append(list, lastCount)
		var cnt int
		for i := 0; i < len(list)-1; i++ {
			cnt += min(list[i], list[i+1])
		}
		return cnt
	*/
	// 方法三：在方法二的基础上，去掉临时数组，因为我们只需要上一个字符的数量
	var (
		one, other int
		result     int
	)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			one++
		} else {
			result += min(one, other)
			other = one
			one = 1
		}
	}
	return result + min(one, other)
}
