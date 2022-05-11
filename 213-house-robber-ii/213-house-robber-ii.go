func rob(nums []int) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        
        return b
    }
    _rob := func(homes []int) int {
        sum0, sum1 := 0, 0
        for _, n := range homes {
            temp := max(n + sum0, sum1)
            sum0 = sum1
            sum1 = temp
        }
        return sum1
    }
    
    if len(nums) == 1 {
        return nums[0]
    }
    
    return max(
        _rob(nums[:len(nums)-1]),
        _rob(nums[1:]),
    )
}

