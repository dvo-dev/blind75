func isAnagram(s string, t string) bool {
    var sTotal int
    sMap := make(map[rune]int)
    for _, r := range s {
        val, ok := sMap[r]
        if ok {
            sMap[r] = val + 1
        } else {
            sMap[r] = 1
        }
        sTotal += 1
    }
    
    for _, r := range t {
        if sTotal == 0 {
            return false
        }
        
        val, ok := sMap[r]
        if !ok || val == 0 {
            return false
        }
        
        sMap[r] = val-1
        sTotal -= 1
    }
    
    return sTotal == 0
}