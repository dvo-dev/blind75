func isPalindrome(s string) bool {
    // Base case
    if len(s) == 0 {
        return true
    }
    
    // Remove spaces, non-alphanumerics, lowercase entire string
    s = (strings.ToLower(
        strings.ReplaceAll(s, " ", ""),
    ))
    reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
    if err != nil {
        log.Fatal(err)
    }
    runes := []rune(reg.ReplaceAllString(s, ""))
    
    // Iterate from both ends of the array, if no mismatches occur, is palindrome
    for left, right := 0, len(runes)-1;
    left < len(runes) && right >= 0;
    left, right = left+1, right-1 {
        if runes[left] != runes[right] {
            return false
        }
    }
    
    return true
}