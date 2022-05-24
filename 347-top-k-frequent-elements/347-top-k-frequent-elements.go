func topKFrequent(nums []int, k int) []int {
    // First determine the frequency of each unique value in the array
    frequency := make(map[int]int)
    for _, n := range nums {
        val, ok := frequency[n]
        if ok {
            frequency[n] = val + 1
        } else {
            frequency[n] = 1
        }
    }
    
    // Rearrange it so we have a slice with index representing frequency,
    // values being a slice with values that occur at the respective freq
    count := make([][]int, len(nums)+1)
    for val, freq := range frequency {
        count[freq] = append(count[freq], val)
    }
    
    // Calculate the result, by traversing backwards from the count slice,
    // until k values are found.
    res := make([]int, k)
    for i := len(count) - 1; i >= 0; i-- {
        for _, n := range count[i] {
            if k-1 < 0 {
                return res
            }

            res[k-1] = n
            k -= 1
        }
    }
    
    return res
}