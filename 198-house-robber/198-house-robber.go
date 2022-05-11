func rob(nums []int) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        
        return b
    }
    var sum0, sum1 int
    
    for _, n := range nums {
        // sum0 is the preceding non adjacent, so you can add n
        // to use sum1 you cannot use the adjacent n
        tmp := max(n + sum0, sum1)
        
        // Advance the pointers, sum1 will be the best value
        sum0 = sum1
        sum1 = tmp
    }
    
    return sum1
}