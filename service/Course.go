package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ljcbaby/plan/config"
	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
	"github.com/tealeg/xlsx"
)

type CourseService struct{}

func (c *CourseService) CreateCourse(course *model.Course) (uint, error) {
	db := database.DB

	if err := db.Create(course).Error; err != nil {
		return 0, err
	}

	return *course.ID, nil
}

func (c *CourseService) DeleteCourse(id uint) error {
	db := database.DB

	if err := db.Delete(&model.Course{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (c *CourseService) GetCourseByCode(code string) (json.RawMessage, error) {
	var course model.Course
	db := database.DB

	if err := db.Where("code = ?", code).First(&course).Error; err != nil {
		return nil, err
	}

	result, err := json.Marshal(course)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CourseService) UpdateCourse(id uint, course *model.Course) error {
	db := database.DB

	var old_course model.Course
	if err := db.Where("id = ?", id).First(&old_course).Error; err != nil {
		return err
	}

	tx := db.Begin()

	if course.Code != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("code", course.Code).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Name != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("name", course.Name).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.ForeignName != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("foreign_name", course.ForeignName).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Credit != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("credit", course.Credit).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.HoursTotal != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_total", course.HoursTotal).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.HoursLecture != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_lecture", course.HoursLecture).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.HoursPractices != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_practices", course.HoursPractices).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.HoursExperiment != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_experiment", course.HoursExperiment).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.HoursComputer != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_computer", course.HoursComputer).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.HoursSelf != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_self", course.HoursSelf).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Assessment != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("assessment", course.Assessment).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.ShowRemark != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("show_remark", course.ShowRemark).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Remark != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("remark", course.Remark).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.DepartmentName != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("department_name", course.DepartmentName).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.LeaderName != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("leader_name", course.LeaderName).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (c *CourseService) GetCourseList(page *model.Page, r *model.Course, courses *[]model.Course) error {
	db := database.DB

	if r.Code != nil {
		db = db.Where("code LIKE ?", "%"+*r.Code+"%")
	}
	if r.Name != nil {
		db = db.Where("name LIKE ?", "%"+*r.Name+"%")
	}
	if r.ForeignName != nil {
		db = db.Where("foreign_name LIKE ?", "%"+*r.ForeignName+"%")
	}
	if r.Remark != nil {
		db = db.Where("remark LIKE ?", "%"+*r.Remark+"%")
	}
	if r.ShowRemark != nil {
		db = db.Where("show_remark LIKE ?", "%"+*r.ShowRemark+"%")
	}
	if r.DepartmentName != nil {
		db = db.Where("department_name LIKE ?", "%"+*r.DepartmentName+"%")
	}
	if r.LeaderName != nil {
		db = db.Where("leader_name LIKE ?", "%"+*r.LeaderName+"%")
	}
	if r.Assessment != nil {
		db = db.Where("assessment = ?", *r.Assessment)
	}
	if r.Credit != nil {
		db = db.Where("credit = ?", *r.Credit)
	}

	if err := db.Model(&model.Course{}).Count(&page.Total).Error; err != nil {
		return err
	}

	if err := db.Offset((page.Current - 1) * page.PageSize).Limit(page.PageSize).Find(courses).Error; err != nil {
		return err
	}

	return nil
}

func (c *CourseService) GetCourseListQuery(page *model.Page, q *string, courses *[]model.Course) error {
	db := database.DB

	db = db.Where("code LIKE ?", "%"+*q+"%").Or("name LIKE ?", "%"+*q+"%").
		Or("foreign_name LIKE ?", "%"+*q+"%").Or("remark LIKE ?", "%"+*q+"%").
		Or("show_remark LIKE ?", "%"+*q+"%")

	if err := db.Model(&model.Course{}).Count(&page.Total).Error; err != nil {
		return err
	}

	if err := db.Offset((page.Current - 1) * page.PageSize).Limit(page.PageSize).Find(courses).Error; err != nil {
		return err
	}

	// for i := 0; i < len(*courses); i++ {
	// 	db := database.DB
	// 	var t sql.NullString
	// 	if err := db.Model(&model.Course{}).Where("id = ?", (*courses)[i].ID).Select("hours_total").Scan(&t).Error; err != nil {
	// 		return err
	// 	}
	// 	(*courses)[i].HoursTotal = new(interface{})
	// 	if t.Valid {
	// 		*(*courses)[i].HoursTotal = t.String
	// 	} else {
	// 		var sum int
	// 		if (*courses)[i].HoursLecture != nil {
	// 			sum += *(*courses)[i].HoursLecture
	// 		}
	// 		if (*courses)[i].HoursPractices != nil {
	// 			sum += *(*courses)[i].HoursPractices
	// 		}
	// 		if (*courses)[i].HoursExperiment != nil {
	// 			sum += *(*courses)[i].HoursExperiment
	// 		}
	// 		if (*courses)[i].HoursComputer != nil {
	// 			sum += *(*courses)[i].HoursComputer
	// 		}
	// 		*(*courses)[i].HoursTotal = sum
	// 	}
	// }

	return nil
}

func (c *CourseService) ReleaseTemplate() error {
	filename := "course-template.xlsx"
	filepath := path.Join(config.Conf.Download.SavePath, filename)

	if _, err := os.Stat(filepath); err != nil {
		db := database.DB

		var template struct {
			Data []byte `gorm:"type:longblob"`
		}
		if err := db.Table("templates").Select("data").Where("name = ?", "courses").
			Scan(&template).Error; err != nil {
			return err
		}

		if len(template.Data) == 0 {
			return errors.New("template data is empty")
		}

		if _, err := os.Stat(config.Conf.Download.SavePath); err != nil {
			if err := os.MkdirAll(config.Conf.Download.SavePath, 0775); err != nil {
				return err
			}
		}

		if err := os.WriteFile(filepath, template.Data, 0775); err != nil {
			return err
		}
	}

	return nil
}

func (c *CourseService) ExportFile() (string, error) {
	filename := "course-" + time.Now().Format("20060102150405") + ".xlsx"
	filepath := path.Join(config.Conf.Download.SavePath, filename)

	if _, err := os.Stat(config.Conf.Download.SavePath); err != nil {
		if err := os.MkdirAll(config.Conf.Download.SavePath, 0775); err != nil {
			return "", err
		}
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet")
	if err != nil {
		return "", err
	}

	header := sheet.AddRow()
	header.AddCell().Value = "课程号"
	header.AddCell().Value = "课程名"
	header.AddCell().Value = "课程外文名"
	header.AddCell().Value = "学分"
	header.AddCell().Value = "总学时"
	header.AddCell().Value = "讲授"
	header.AddCell().Value = "课程实践"
	header.AddCell().Value = "实验"
	header.AddCell().Value = "课内上机"
	header.AddCell().Value = "课外上机"
	header.AddCell().Value = "考核方式"
	header.AddCell().Value = "展示备注"
	header.AddCell().Value = "开课备注"
	header.AddCell().Value = "开课学院"
	header.AddCell().Value = "课程负责人"

	var courses []model.Course
	err = c.GetCourseList(&model.Page{Current: 1, PageSize: 100000}, &model.Course{}, &courses)
	if err != nil {
		return "", err
	}

	for _, course := range courses {
		row := sheet.AddRow()
		row.AddCell().Value = *course.Code
		row.AddCell().Value = *course.Name
		row.AddCell().Value = *course.ForeignName
		row.AddCell().Value = fmt.Sprintf("%v", *course.Credit)
		row.AddCell().Value = fmt.Sprintf("%v", *course.HoursTotal)
		if course.HoursLecture != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.HoursLecture)
		} else {
			row.AddCell().Value = ""
		}
		if course.HoursPractices != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.HoursPractices)
		} else {
			row.AddCell().Value = ""
		}
		if course.HoursExperiment != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.HoursExperiment)
		} else {
			row.AddCell().Value = ""
		}
		if course.HoursComputer != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.HoursComputer)
		} else {
			row.AddCell().Value = ""
		}
		if course.HoursSelf != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.HoursSelf)
		} else {
			row.AddCell().Value = ""
		}
		row.AddCell().Value = *course.Assessment
		if course.ShowRemark != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.ShowRemark)
		} else {
			row.AddCell().Value = ""
		}
		if course.Remark != nil {
			row.AddCell().Value = fmt.Sprintf("%v", *course.Remark)
		} else {
			row.AddCell().Value = ""
		}
		row.AddCell().Value = *course.DepartmentName
		row.AddCell().Value = *course.LeaderName
	}

	if err := file.Save(filepath); err != nil {
		return "", err
	}

	return filename, nil
}

func (c *CourseService) ImportFile(file []byte, sun *uint, errs *[]string) error {
	xlFile, err := xlsx.OpenBinary(file)
	if err != nil {
		return err
	}

	if len(xlFile.Sheets) == 0 {
		return errors.New("sheet is empty")
	}

	db := database.DB

	sheet := xlFile.Sheets[0]
	tx := db.Begin()
	for i, row := range sheet.Rows {
		if i == 0 {
			continue
		}

		var course model.Course
		for j, cell := range row.Cells {
			switch j {
			case 0:
				course.Code = new(string)
				*course.Code = cell.String()
				if *course.Code == "" {
					course.Code = nil
				}
			case 1:
				course.Name = new(string)
				*course.Name = cell.String()
				if *course.Name == "" {
					course.Name = nil
				}
			case 2:
				course.ForeignName = new(string)
				*course.ForeignName = cell.String()
				if *course.ForeignName == "" {
					course.ForeignName = nil
				}
			case 3:
				course.Credit = new(float64)
				*course.Credit, err = cell.Float()
				if err != nil {
					course.Credit = nil
				}
			case 4:
				course.HoursTotal = new(string)
				*course.HoursTotal = cell.String()
				if *course.HoursTotal == "" {
					course.HoursTotal = nil
				}
			case 5:
				course.HoursLecture = new(int)
				*course.HoursLecture, err = cell.Int()
				if err != nil {
					course.HoursLecture = nil
				}
			case 6:
				course.HoursPractices = new(int)
				*course.HoursPractices, err = cell.Int()
				if err != nil {
					course.HoursPractices = nil
				}
			case 7:
				course.HoursExperiment = new(int)
				*course.HoursExperiment, err = cell.Int()
				if err != nil {
					course.HoursExperiment = nil
				}
			case 8:
				course.HoursComputer = new(int)
				*course.HoursComputer, err = cell.Int()
				if err != nil {
					course.HoursComputer = nil
				}
			case 9:
				course.HoursSelf = new(int)
				*course.HoursSelf, err = cell.Int()
				if err != nil {
					course.HoursSelf = nil
				}
			case 10:
				course.Assessment = new(string)
				*course.Assessment = cell.String()
				if *course.Assessment == "" {
					course.Assessment = nil
				}
			case 11:
				course.ShowRemark = new(string)
				*course.ShowRemark = cell.String()
				if *course.ShowRemark == "" {
					course.ShowRemark = nil
				}
			case 12:
				course.Remark = new(string)
				*course.Remark = cell.String()
				if *course.Remark == "" {
					course.Remark = nil
				}
			case 13:
				course.DepartmentName = new(string)
				*course.DepartmentName = cell.String()
				if *course.DepartmentName == "" {
					course.DepartmentName = nil
				}
			case 14:
				course.LeaderName = new(string)
				*course.LeaderName = cell.String()
				if *course.LeaderName == "" {
					course.LeaderName = nil
				}
			}
		}

		if course.Code == nil || course.Name == nil || course.ForeignName == nil || course.Credit == nil ||
			course.HoursTotal == nil || course.Assessment == nil || course.DepartmentName == nil ||
			course.LeaderName == nil {
			*errs = append(*errs, "L"+fmt.Sprintf("%d", i+1)+": 必填字段不能为空")
			continue
		}

		switch *course.Assessment {
		case "X", "Y", "C":
		default:
			*errs = append(*errs, "L"+fmt.Sprintf("%d", i+1)+": 考核方式不正确")
			continue
		}

		if err := tx.Create(&course).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				*errs = append(*errs, "L"+fmt.Sprintf("%d", i+1)+": 课程号已存在")
				continue
			} else {
				tx.Rollback()
				return err
			}
		} else {
			*sun++
		}
	}

	tx.Commit()

	return nil
}
