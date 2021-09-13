package lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	first, second, ans := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 {
			second = i
			ans = max(ans, second-first+1)
		} else {
			second = i - 1
			ans = max(second-first+1, ans)
			temp := m[s[i]]
			for k := first; k < temp; k++ {
				delete(m, s[k])
			}
			first = temp
		}
		m[s[i]] = i +1
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
