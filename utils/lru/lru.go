package lru

import (
	"errors"
	"fmt"
)

//LRU实现思路
//我们维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。当有一个新的数据被访问时，我们从链表头开始顺序遍历链表。
//
//如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
//如果此数据没有在缓存链表中，又可以分为两种情况：
//如果此时缓存未满，则将此结点直接插入到链表的头部；
//如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部。
//补充：现在我们来看下 m 缓存访问的时间复杂度是多少。因为不管缓存有没有满，我们都需要遍历一遍链表，所以这种基于链表的实现思路，缓存访问的时间复杂度为 O(n)。实际上，我们可以继续优化这个实现思路，比如引入散列表（Hash table）来记录每个数据的位置，将缓存访问的时间复杂度降到 O(1)。另外除了基于链表的实现思路，实际上还可以用数组来实现 LRU 缓存淘汰策略。

//我们维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。当有一个新的数据被访问时，我们从链表头开始顺序遍历链表。
//
//如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
//如果此数据没有在缓存链表中，又可以分为两种情况：
//如果此时缓存未满，则将此结点直接插入到链表的头部；
//如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部。

type Node struct {
	data int
	next *Node
}

type HeadNode struct {
	num   int
	first *Node
}

var (
	LIST_LEN = 5
)

//打印
func printNode(head *HeadNode) error {
	if head == nil {
		return errors.New("input parameters is nil")
	}

	if head.first == nil {
		fmt.Printf("node is null\n")
	} else {
		curNode := head.first

		for curNode != nil {
			fmt.Printf("data is: %v\n", curNode)
			curNode = curNode.next
		}
	}

	return nil
}

//尾插入
func insertTailNode(head *HeadNode, node *Node) error {
	if head == nil || node == nil {
		return errors.New("input parameters is nil")
	}

	if head.first == nil {
		head.num = 1
		head.first = node
	} else {
		curNode := head.first

		for curNode.next != nil {
			curNode = curNode.next
		}

		curNode.next = node
		head.num++
	}
	return nil
}

//头插入
func insertHeadNode(head *HeadNode, node *Node) error {
	if head == nil || node == nil {
		return errors.New("input parameters is nil")
	}

	if head.first == nil {
		head.num = 1
		head.first = node
	} else {
		curNode := head.first
		head.first = node
		node.next = curNode

		head.num++
	}

	return nil
}

//判断节点是否存在列表，存在则返回该节点，不存在则返回nil
func selectNode(head *HeadNode, node *Node) (*Node, error) {
	if head == nil || node == nil {
		return nil, errors.New("input parameters is nil")
	}

	if head.first == nil {
		return nil, nil
	} else {
		curNode := head.first

		for curNode != nil {
			if node.data == curNode.data {
				return curNode, nil
			}
			curNode = curNode.next
		}
	}

	return nil, nil
}

//删除节点
func deleteNode(head *HeadNode, node *Node) error {
	if head == nil || node == nil {
		return errors.New("input parameters is nil")
	}

	count := 0

	if head.first == nil {
		return errors.New("list is nil")
	} else {
		curNode, PreNode := head.first, head.first

		for curNode != nil {
			count++
			if node.data == curNode.data {
				// 如果仅有一个节点，则头尾节点清空
				if head.num == 1 {
					head.first = nil
					head.num--
					return nil
				}
				// 若删除的尾部，尾部指针需要调整
				if head.num == count {
					PreNode.next = nil
					head.num--
					return nil
				}

				PreNode.next = PreNode.next.next
				head.num--
				return nil

			}
			PreNode = curNode
			curNode = curNode.next
		}
	}

	return nil
}

//删除尾节点
func deleteTailNode(head *HeadNode) error {
	if head == nil {
		return errors.New("input parameters is nil")
	}

	if head.first == nil {
		return errors.New("list is nil")
	} else {
		curNode := head.first

		// 如果仅有一个节点，则头尾节点清空
		if head.num == 1 {
			head.first = nil
			head.num--
			return nil
		}

		for {
			if curNode.next != nil && curNode.next.next != nil {
				curNode = curNode.next
			} else {
				break
			}
		}
		//倒数第二个结点的next赋nil
		curNode.next = nil
		head.num--
	}

	return nil
}

//处理缓存淘汰列表
func handleLRU(head *HeadNode, node *Node) error {
	if head == nil || node == nil {
		return errors.New("input parameters is nil")
	}

	tmpNode, err := selectNode(head, node)
	if err != nil {
		return err
	}

	if tmpNode == nil {
		//如果此数据没有在缓存链表中 如果大于缓存长度10，则删除尾巴节点
		if head.num >= LIST_LEN {
			err = deleteTailNode(head)
			if err != nil {
				return err
			}
		}

		err = insertHeadNode(head, node)
		if err != nil {
			return err
		}

	} else {
		//如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
		err = deleteNode(head, tmpNode)
		if err != nil {
			return err
		}

		err = insertHeadNode(head, tmpNode)
		if err != nil {
			return err
		}

	}

	return nil
}
