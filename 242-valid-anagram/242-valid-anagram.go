func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sMap, tMap := make(map[rune]int), make(map[rune]int)
    sRunes, tRunes := []rune(s), []rune(t)
    
    for i := range sRunes {
        val, ok := sMap[sRunes[i]]
        if ok {
            sMap[sRunes[i]] = val + 1
        } else {
            sMap[sRunes[i]] = 1
        }
        
        val, ok = tMap[tRunes[i]]
        if ok {
            tMap[tRunes[i]] = val + 1
        } else {
            tMap[tRunes[i]] = 1
        }
    }
    
    for k, v := range sMap {
        if val, ok := tMap[k]; !ok || v != val {
            return false
        }
    }
    
    return len(sMap) == len(tMap)
}