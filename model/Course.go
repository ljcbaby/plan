package model

// 课程
type Course struct {
	ID              uint       `gorm:"primaryKey" json:"id,omitempty"`        // 课程号
	Name            string     `gorm:"index" json:"name,omitempty"`           // 课程名
	ForeignName     string     `gorm:"index" json:"foreignName,omitempty"`    // 课程外文名
	Credit          int        `json:"credit,omitempty"`                      // 学分
	HoursTotal      HoursTotal `json:"hoursTotal"`                            // 总学时
	HoursLecture    int        `json:"hoursLecture,omitempty"`                // 讲授
	HoursPractices  int        `json:"hoursPractices,omitempty"`              // 课程实践
	HoursExperiment int        `json:"hoursExperiment,omitempty"`             // 实验
	HoursComputer   int        `json:"hoursComputer,omitempty"`               // 课内上机
	HoursSelf       int        `json:"hoursSelf,omitempty"`                   // 课外上机
	Assessment      string     `gorm:"index" json:"assessment,omitempty"`     // 考核方式，X代表“学校组织考试”，Y代表“学院组织考试”，C代表“考查”。
	ShowRemark      string     `gorm:"index" json:"showRemark,omitempty"`     // 展示备注，双语、全英文 等
	Remark          string     `json:"remark,omitempty"`                      // 开课备注
	DepartmentName  string     `gorm:"index" json:"departmentName,omitempty"` // 开课学院
	LeaderName      string     `gorm:"index" json:"leaderName,omitempty"`     // 课程负责人
}

type HoursTotal struct {
	Integer *int64  `json:"integer,omitempty"`
	String  *string `json:"string,omitempty"`
}
