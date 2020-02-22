package leetcode

/*
  给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			switch {
			case i == 0 && j == 0:
				dp[j] = true
			case i == 0:
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1]
			case j == 0:
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			default:
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1] ||
					dp[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return dp[len(s2)]
}