package ecs

// 基础组件
type Component struct {
	Parent *Entity // 组件所挂载的实体
}
