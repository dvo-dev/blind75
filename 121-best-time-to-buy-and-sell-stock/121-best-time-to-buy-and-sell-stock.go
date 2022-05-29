func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    
    buy, profit := 0, 0
    for i, p := range prices {
        // Always attempt to buy at lowest value
        if prices[buy] > p {
            buy = i
        }
        
        // Calculate if this new profit is actually the new highest
        profit = max(profit, p - prices[buy])
    }
    
    return profit
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}
