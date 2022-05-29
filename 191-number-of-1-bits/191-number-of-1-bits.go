func hammingWeight(num uint32) int {
    var count int
    for num != 0 {
        // LSB check
        if num & 1 != 0 {
            count++
        }
        
        // "advance" the value to check next
        num = num>>1
    }
    
    return count
}