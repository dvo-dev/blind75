func climbStairs(n int) int {
    // Base cases
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if n == 2{
        return 2
    }
    
    // Memoization
    dp := make([]int, n+1)
    dp[n], dp[n-1] = 1, 1
    
    // Reuse memorized values
    for i := n-2; i >= 0; i-- {
        dp[i] = dp[i+1] + dp[i+2]
    }
    
    return dp[0]
}

func bruteForceClimbStairs(n int) int {
    if n - 1 < 0 {
        return 0
    }
    
    if n - 1 == 0 {
        return 1
    }
    
    return climbStairs(n - 1) + func() int {
        if n - 2 < 0 {
            return 0
        } else if n - 2 == 0 {
            return 1
        } else {
            return climbStairs(n - 2)
        }
    }()
}