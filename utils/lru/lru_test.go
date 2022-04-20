package lru

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	head := &HeadNode{0, nil}

	//缓存长度为5
	//case 1 如果此数据没有在缓存链表中，且链表已满，则删除尾节点
	_ = handleLRU(head, &Node{1, nil})
	_ = handleLRU(head, &Node{2, nil})
	_ = handleLRU(head, &Node{3, nil})
	_ = handleLRU(head, &Node{4, nil})
	_ = handleLRU(head, &Node{5, nil})
	_ = handleLRU(head, &Node{6, nil})
	_ = printNode(head)
	fmt.Printf("all node num is: %v\n\n", head.num)

	//case 2 如果此数据没有在缓存链表中，且链表未满，则将此结点直接插入到链表的头部
	_ = deleteNode(head, &Node{4, nil})
	_ = deleteNode(head, &Node{3, nil})
	_ = printNode(head)
	fmt.Printf("after delete list num is: %v\n\n", head.num)
	_ = handleLRU(head, &Node{7, nil})
	_ = printNode(head)
	fmt.Printf("all node num is: %v\n\n", head.num)

	//case 3 如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
	_ = insertHeadNode(head, &Node{8, nil})
	_ = printNode(head)
	fmt.Printf("after add node list num is: %v\n\n", head.num)
	_ = handleLRU(head, &Node{5, nil})
	_ = printNode(head)
	fmt.Printf("all node num is: %v\n\n", head.num)

}
