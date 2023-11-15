package model

import "encoding/json"

// 课程
type Course struct {
	ID              *uint        `gorm:"primaryKey" json:"id,omitempty"`
	Code            *string      `json:"code,omitempty"`                          // 课程号
	Name            *string      `gorm:"index" json:"name,omitempty"`             // 课程名
	ForeignName     *string      `gorm:"index" json:"foreignName,omitempty"`      // 课程外文名
	Credit          *float64     `json:"credit,omitempty"`                        // 学分
	HoursTotal      *interface{} `gorm:"type:string" json:"hoursTotal,omitempty"` // 总学时
	HoursLecture    *int         `json:"hoursLecture,omitempty"`                  // 讲授
	HoursPractices  *int         `json:"hoursPractices,omitempty"`                // 课程实践
	HoursExperiment *int         `json:"hoursExperiment,omitempty"`               // 实验
	HoursComputer   *int         `json:"hoursComputer,omitempty"`                 // 课内上机
	HoursSelf       *int         `json:"hoursSelf,omitempty"`                     // 课外上机
	Assessment      *string      `gorm:"index" json:"assessment,omitempty"`       // 考核方式，X代表“学校组织考试”，Y代表“学院组织考试”，C代表“考查”。
	ShowRemark      *string      `gorm:"index" json:"showRemark,omitempty"`       // 展示备注，双语、全英文 等
	Remark          *string      `json:"remark,omitempty"`                        // 开课备注
	DepartmentName  *string      `gorm:"index" json:"departmentName,omitempty"`   // 开课学院
	LeaderName      *string      `gorm:"index" json:"leaderName,omitempty"`       // 课程负责人
}

func (c *Course) UnmarshalJSON(data []byte) error {
	type Alias Course
	aux := &struct {
		HoursTotal json.RawMessage `json:"hoursTotal,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if len(aux.HoursTotal) > 0 {
		if aux.HoursTotal[0] == '"' {
			var hoursTotal string
			if err := json.Unmarshal(aux.HoursTotal, &hoursTotal); err != nil {
				return err
			}
			*c.HoursTotal = hoursTotal
		} else {
			var hoursTotal int
			if err := json.Unmarshal(aux.HoursTotal, &hoursTotal); err != nil {
				return err
			}
			*c.HoursTotal = hoursTotal
		}
	} else {
		*c.HoursTotal = nil
	}

	return nil
}
