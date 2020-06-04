package mylib
//package main

import "fmt"

// ====================================================================
// struct defination
// ====================================================================


type Node struct{
	left *Node
	data interface{}
	right *Node
}


// ====================================================================
// API defination
// ====================================================================

func NewNode(left, right *Node) *Node {
    return &Node{left, nil, right}
}

func (n *Node)SetData(data interface{}){
	n.data=data
}

// 			1
//    	   / \  
//        2   3
//       / \   \
//      4   5   6
//         / \
//        7   8

//	前序遍历：1  2  4  5  7  8  3  6
//	后序遍历：4  7  8  5  2  6  3  1 
//	中序遍历：4  2  7  5  8  1  3  6

func TestNode(){
	One := NewNode(nil,nil)
	One.SetData(1)	
	Two := NewNode(nil,nil)
	Two.SetData(2)
	Three := NewNode(nil,nil)
	Three.SetData(3)
	Four := NewNode(nil,nil)
	Four.SetData(4)
	Five := NewNode(nil,nil)
	Five.SetData(5)
	Six := NewNode(nil,nil)
	Six.SetData(6)
	Seven := NewNode(nil,nil)
	Seven.SetData(7)
	Eight := NewNode(nil,nil)
	Eight.SetData(8)

	One.left=Two
	One.right=Three

	Three.right=Six

	Two.left=Four
	Two.right=Five

	Five.left=Seven
	Five.right=Eight

	Root := One
	fmt.Printf("pre order...	")
	pre_order_traversal(Root)
	fmt.Println()
	fmt.Printf("post order...	")
	post_order_traversal(Root)
	fmt.Println()
	fmt.Printf("in order...		")
	in_order_traversal(Root)
	fmt.Println()
}

//前序,递归调用
func pre_order_traversal(root *Node){
	if root == nil{
		return
	}
	fmt.Printf("%v  ",root.data)
	if root.left != nil{
		pre_order_traversal(root.left)
	}
	if root.right != nil{
		pre_order_traversal(root.right)
	}
}

//前序,递归调用
func pre_order_traversal2(root *Node){
	if root == nil{
		return
	}
	fmt.Printf("%v  ",root.data)
	if root.left != nil{
		pre_order_traversal(root.left)
	}
	if root.right != nil{
		pre_order_traversal(root.right)
	}
}

//后续,递归调用
func post_order_traversal(root *Node){
	if root == nil{
		return
	}	
	if root.left != nil{
		post_order_traversal(root.left)
	}
	if root.right != nil{
		post_order_traversal(root.right)
	}
	fmt.Printf("%v  ",root.data)
}

//中序,递归调用
func in_order_traversal(root *Node){
	if root == nil{
		return
	}
	if root.left != nil{
		in_order_traversal(root.left)
	}
	fmt.Printf("%v  ",root.data)
	if root.right != nil{
		in_order_traversal(root.right)
	}
}