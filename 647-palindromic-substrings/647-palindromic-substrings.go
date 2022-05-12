func countSubstrings(s string) int { 
    var count int
    for i := range s {
        count += countPalindrome(i, i, s)   // Using current as centerpiece (odd palindromes)
        count += countPalindrome(i, i+1, s) // Use adjacents as centerpiece (even palindromes)
    }
    
    return count
}

func countPalindrome(left, right int, s string) int {
    // Block out of bounds
    if left < 0 || right >= len(s) {
        return 0
    }
    var count int
    
    // Expand outwards from the two pointers while
    // checking for a match
    for left, right = left, right;
    left >= 0 && right < len(s) && s[left] == s[right]; 
    left, right = left - 1, right +1 {
        count += 1
    }
    
    // Return number of palindromes found
    return count
}