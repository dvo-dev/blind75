func twoSum(nums []int, target int) []int {
    // Impossible edge case
    if len(nums) < 2 {
        return []int{}
    }
    
    // Track numbers and indexes
    set := make(map[int]int)
    
    // Range through, store the required value to reach target
    // for each value - succeeds when the current iteration can
    // be combined with an existing value to reach target
    for i, n := range nums {
        if idx, ok := set[n]; ok {
            return []int{i, idx}
        }
        
        set[target-n] = i
    }
    
    return []int{}
}