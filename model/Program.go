package model

// 培养计划
type Program struct {
	ID           *uint     `gorm:"primaryKey" json:"id,omitempty"` // ID
	Name         *string   `json:"name,omitempty"`                 // 计划名
	Major        *string   `json:"major,omitempty"`                // 专业
	Department   *string   `json:"department,omitempty"`           // 描述
	DependencyID *uint     `json:"dependencyId,omitempty"`         // 依赖ID
	Grade        *string   `json:"grade,omitempty"`                // 年级
	Content      *[]Node   `json:"content,omitempty"`              // 内容
	Tags         *[]string `gorm:"-" json:"tags,omitempty"`        // 标签
}

// 节点
type Node struct {
	ID          *uint        `json:"id,omitempty"`          // 节点ID
	Name        *string      `json:"name,omitempty"`        // 节点名称
	BeginWeek   *string      `json:"beginWeek,omitempty"`   // 起始周
	CourseID    *string      `json:"courseId,omitempty"`    // 课程号
	Remark      *string      `json:"remark,omitempty"`      // 备注
	Semester    *string      `json:"semester,omitempty"`    // 学期
	Tags        *[]string    `json:"tags,omitempty"`        // 标签
	Type        *string      `json:"type,omitempty"`        // 类型，node or course
	Requirement *Requirement `json:"requirement,omitempty"` // 要求
	Content     *[]Node      `json:"content,omitempty"`     // 包含节点/课程
	ParentID    *uint        `json:"parentId,omitempty"`    // 父节点ID
}

// 要求
type Requirement struct {
	ID        uint  `json:"id,omitempty"`        // 要求ID
	MinCourse int64 `json:"minCourse,omitempty"` // 最少修习门数
	MinCredit int64 `json:"minCredit,omitempty"` // 最低学分
	NodeID    uint  `json:"nodeId,omitempty"`    // 对应节点ID
}
