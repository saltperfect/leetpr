package dp

import "fmt"

/*
	 ______________
	|d ||  ||  ||  ||  ||  ||  ||  |		|  |
	|__||__||__||__||__||__||__||__|		|__|
	|  ||  ||  ||  ||  ||  ||  ||  |		|  |
	|__||__||__||__||__||__||__||__|		|__|
	|  ||  ||  ||  ||  ||  ||  ||  |		|  |
	|__||__||__||__||__||__||__||__|		|__|
	|  ||  ||  ||  ||  ||  ||  ||  |		|  |
	|__||__||__||__||__||__||__||__|		|__|
	|  ||  ||  ||  ||  ||  ||  ||  |		|  |
	|__||__||__||__||__||__||__||__|		|__|
*/

// dp[pattern][string]
func isMatch(s string, p string) bool {
	r := len(p)
	c := len(s)

	dp := make([][]bool, r+1)
	for i := range dp {
		dp[i] = make([]bool, c+1)
	}

	dp[0][0] = true

	for i := 2; i <= r; i += 2 {
		if p[i-1] == '*' && dp[i-2][0] {
			dp[i][0] = true
		}
	}


	for i := 1; i <= r; i++ {
		for j := 1; j <=c; j++ {
			cs, cp := s[j-1], p[i-1]
			if cs == cp || cp == '.'{
				dp[i][j] = dp[i-1][j-1]
			}else if cp == '*' {
				precp := p[i-2]
				if precp != '.' && precp != cs {
					dp[i][j] = dp[i-2][j]
				}else {
					dp[i][j] = dp[i-2][j] || dp[i-2][j-1] || dp[i][j-1]
				}
			}
		}
	}
	for i := 0; i <= r; i++ {
		fmt.Println(dp[i])
	}
	return dp[r][c]

}
