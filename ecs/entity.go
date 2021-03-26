package ecs

import (
	"math"
)

// 实体基础
type Entity struct {
	Id       uint64            // 实体ID
	Parent   *BasicEntity      // 父节点
	Children []*BasicEntity    // 子节点
	Cmpts    map[string]uint64 // 组件类型 -> 实例 映射
}

// 添加一个子节点
func (this *Entity) AddChild(child *Entity) {
	if child != nil {
		this.Parent = child
		this.Children = append(this.Children, child)
	}
}

// 移除一个子节点
func (this *Entity) RemoveChild(child *Entity) {
	if child == nil {
		return
	}

	index := -1
	for i, e := range this.Children {
		if e.Id == child.Id {
			index = i
			break
		}
	}

	if index >= 0 {
		this.children = append(this.children[:index], this.children[index+1:]...)
	}
}

// 添加组件
func (this *Entity) AddComponent(cmpt Component) {
	// 通过管理系统创建一个 cmpt 实例
	// 检查是否拥有相同类型的组件
	// 将组件添加至 Map 列表
}

// 移除组件
func (this *Entity) RemoveComponent(cmpt Component) {
	// 通过管理系统移除组件
	// 实体系统移除组件
}
