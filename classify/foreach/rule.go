package foreach

// 比特与 2 比特字符
// 有两种特殊字符：
// 第一种字符可以用一比特0 表示
// 第二种字符可以用两比特（10或11）表示
// 给你一个以 0 结尾的二进制数组bits，如果最后一个字符必须是一个一比特字符，则返回 true 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/1-bit-and-2-bit-characters
func isOneBitCharacter(bits []int) bool {
	// 如果碰到1，则必是两比特字符，走两步，如果碰到1，则必是1比特字符，走一步。最后判断前n-1个比特是否正常走完
	n := len(bits)
	i := 0
	for i < n-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == n-1
}

// 推多米诺
// n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
// 每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
// 如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
// 就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
// 给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
// dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
// dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
// dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
// 返回表示最终状态的字符串。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/push-dominoes
func pushDominoes(dominoes string) string {
	// 记录当前牌被推的方向
	preIsR := false // 只有前边的方向是向右推才对后边的牌有效
	num := 0
	n := len(dominoes)
	result := []byte(dominoes)
	for i := 0; i < n; i++ {
		cur := dominoes[i]
		if cur == '.' {
			num++
			continue
		}
		if cur == 'R' {
			// R..R或者...R
			// 被向右推时，处理前n张牌。此时cur对前n张牌不起作用
			// 如果前边有向右推的力，则表示为向右；否则维持原状
			if preIsR {
				for j := i - num; j < i; j++ {
					result[j] = 'R'
				}
			}
			num, preIsR = 1, true
		} else if cur == 'L' {
			// R..L 或者 ...L
			// 如果前边有向右推的力，则需要进行“平衡”；否则，都倒向左
			if preIsR {
				for l, r := i-num, i; l < r; l, r = l+1, r-1 {
					result[l], result[r] = 'R', 'L'
				}
			} else {
				for j := i - num; j <= i; j++ {
					result[j] = 'L'
				}
			}
			num, preIsR = 0, false
		}
	}
	// 处理最后一个方向是向右的情况
	if preIsR && num > 1 {
		for j := n - num; j < n; j++ {
			result[j] = 'R'
		}
	}
	return string(result)
}
