package main

import (
	"LearnGo/container/tree"
	"fmt"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 10:10
 * @Version 1.0
 */

//通过组合的方式对已有类型进行拓展
type myTreeNode struct {
	node *tree.TreeNode
}

//通过别名的方式

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}
func main() {

	var root tree.TreeNode
	fmt.Println(root)
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	//不论是地址还是结构本身，一律使用。来访问成员
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
	root.Print()
	root.Right.Left.SetValue(4)
	fmt.Println(root.Right.Left.Value)

	pRoot := &root
	pRoot.SetValue(5)
	pRoot.Print()
	//root的值被改变了
	root.Print()

	//nil也可以进入函数 运行
	var pRoot2 *tree.TreeNode
	pRoot2.SetValue(200)
	pRoot2 = &root
	pRoot2.SetValue(300)
	pRoot2.Print()
	//实现遍历函数
	fmt.Println()
	root.Traverse()
	/*要改变内容必须使用指针接收者
	结构也大也使用指针接收者
	一致性，如果有指针接收者，最好都是指针接收者
	*/
	fmt.Println()
	myroot := myTreeNode{&root}
	myroot.postOrder()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.TreeNode) {
		nodeCount++
	})
	fmt.Println("NodeCount:", nodeCount)

}
