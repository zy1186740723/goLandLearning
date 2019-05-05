package tree

import "fmt"

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 9:16
 * @Version 1.0
 */
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//为结构定义方法,接收者的概念，print方法是被node所接受的
//值接收者
func (node TreeNode) Print() {
	fmt.Print(node.Value)
}

//指针接收者
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = value
}

//使用工厂函数返回结构的地址，是局部变量的地址
//不需要知道分配在堆还是栈上 依据情况分析
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
