package queue

/**如何拓展已有的数据结构，分装 定义别名和使用组合
 * @Author: zhangyan
 * @Date: 2019/5/5 10:47
 * @Version 1.0
 */
//对silice进行别名
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	//对类型进行限定
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0

}
