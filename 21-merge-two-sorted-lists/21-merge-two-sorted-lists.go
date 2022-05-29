/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := new(ListNode)  // Dummy tracker
    tail := head    // Actual final LL iterator
    
    for list1 != nil && list2 != nil {
        // Add next ascending value to cumulative LL
        if list1.Val <= list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        
        // Advance to the newly added node
        tail = tail.Next
    }
    
    // Append the remainder if the 2 original LL if applicable
    if list1 != nil && list2 == nil {
        tail.Next = list1
    } else if list1 == nil && list2 != nil {
        tail.Next = list2
    }
    
    return head.Next
}