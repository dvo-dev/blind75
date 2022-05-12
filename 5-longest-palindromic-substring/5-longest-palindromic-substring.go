func longestPalindrome(s string) string {
    maxStr := func(a, b string) string {
        if len(a) > len(b) {
            return a
        }
        
        return b
    }
    
    var resStr string
    for i := range s {
        resStr = maxStr(
            resStr,
            maxStr(
                checkPalindrome(i, i, s),   // Using current as expansion point (odd palindromes)
                checkPalindrome(i, i+1, s), // Using "adjacents" as expansion points (even palindromes)
            ),
        )
    }
    
    return resStr
}

func checkPalindrome(left, right int, s string) string {
    // Block out of bounds
    if left < 0 || right >= len(s) {
        return ""
    }
    
    // Expand outwards from the two pointers while
    // checking for a match
    for left, right = left, right;
    left >= 0 && right < len(s) && s[left] == s[right]; 
    left, right = left - 1, right +1 {}
    
    // Return the longest palindrome
    return s[left+1:right]
}