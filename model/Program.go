package model

import (
	"encoding/json"
)

// 培养计划
type Program struct {
	ID           *uint            `gorm:"primaryKey" json:"id,omitempty"`     // ID
	Name         *string          `json:"name,omitempty"`                     // 计划名
	Major        *string          `json:"major,omitempty"`                    // 专业
	Department   *string          `json:"department,omitempty"`               // 描述
	DependencyID *uint            `json:"dependencyId,omitempty"`             // 依赖ID
	Grade        *string          `json:"grade,omitempty"`                    // 年级
	Content      *json.RawMessage `gorm:"type:json" json:"content,omitempty"` // 内容
	Tags         *[]string        `gorm:"-" json:"tags,omitempty"`            // 标签
}

// 节点
type Node struct {
	ID      *string `json:"id,omitempty"` // 节点ID
	Title   *Title  `json:"title,omitempty"`
	Content *[]Node `json:"content,omitempty"`
}

// 节点信息
type Title struct {
	Tags *[]string `json:"tags,omitempty"` // 标签
	Type *string   `json:"type,omitempty"` // 类型
	TitleNode
	TitleCourse
}

// 节点信息-节点
type TitleNode struct {
	Name        *string      `json:"name,omitempty"`        // 节点名
	Requirement *Requirement `json:"requirement,omitempty"` // 节点要求
	Remark      *string      `json:"remark,omitempty"`      // 备注
	AllCredit   *float64     `json:"allCredit,omitempty"`   // 总学分
}

// 节点信息-课程
type TitleCourse struct {
	CourseCode *string          `json:"courseCode,omitempty"` // 课程号
	Semester   *string          `json:"semester,omitempty"`   // 学期
	BeginWeek  *string          `json:"beginWeek,omitempty"`  // 起始周
	Course     *json.RawMessage `json:"course,omitempty"`     // 课程
}

// 要求
type Requirement struct {
	MinCredit *float64 `json:"minCredit,omitempty"` // 最低学分
	MinCourse *int     `json:"minCourse,omitempty"` // 最低满足节点数
}
