/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    
    // head = current node
    // prev = used to remember the preceeding node to "reverse" the next pointer
    // next = used to keep track of dangling next, so we can advance the loop
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    
    return prev
}