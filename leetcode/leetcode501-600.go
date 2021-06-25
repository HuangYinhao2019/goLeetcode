package leetcode

//516. 最长回文子序列
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i + j < len(s); i++ {
			if s[i] == s[i + j] {
				dp[i][i + j] = dp[i + 1][i + j - 1] + 2
			}
			dp[i][i + j] = max(dp[i][i + j], max(dp[i + 1][i + j], dp[i][i + j - 1]))
		}
	}
	return dp[0][len(s) - 1]
}







