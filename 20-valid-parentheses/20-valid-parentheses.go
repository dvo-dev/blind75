func isValid(s string) bool {
    brackets := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }
    
    st := new(Stack)
    for _, r := range s {
        b, ok := brackets[r]
        if ok {
            st.Push(b)
        } else if st.IsEmpty() || st.Pop() != r {
            fmt.Printf("%v %v", st, r)
            return false
        }
    }
    
    return st.IsEmpty()
}


type Stack []rune

func (s *Stack) Push(r rune) {
    *s = append(*s, r)
}

func (s *Stack) Pop() rune {
    r := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return r
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}