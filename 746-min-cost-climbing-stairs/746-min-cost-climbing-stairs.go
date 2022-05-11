func minCostClimbingStairs(cost []int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        
        return b
    }
    l := len(cost)
    
    // Memoization
    dp := make([]int, len(cost))
    dp[l-1], dp[l-2] = cost[l-1], cost[l-2]  // Base cases
    
    for i := l-3; i >= 0; i-- {
        // Memorizing cost totals from top down
        // cost = cumulative sum of min between 1/2 steps
        dp[i] = cost[i] + min(dp[i+1], dp[i+2])
    }
    
    return min(dp[0], dp[1])
}