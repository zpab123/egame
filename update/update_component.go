package update

type UpdateFunction func(dt)

// 帧函数组件
type UpdateCompoent struct {
	Update UpdateFunction // 帧函数具体的函数
}
