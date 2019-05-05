package tree

import "fmt"

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 10:18
 * @Version 1.0
 */
func (node *TreeNode) Traverse() {
	//if node==nil {
	//	return
	//}
	//node.Left.Traverse()//即使是nil，函数也可以进行，
	////nil只支持指针接收者
	//node.Print()
	//node.Right.Traverse()
	node.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
	fmt.Println()
}

func (node *TreeNode) TraverseFunc(
	f func(node *TreeNode)) {
	if node == nil {
		return
	}
	//函数编程的思想，可传入各种实现的方式f
	node.Left.TraverseFunc(f) //即使是nil，函数也可以进行，
	//nil只支持指针接收者
	f(node)
	node.Right.TraverseFunc(f)

}
