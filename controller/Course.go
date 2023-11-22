package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/model"
	"github.com/ljcbaby/plan/service"
)

type CourseController struct{}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course model.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	if course.Code == nil || course.Name == nil || course.ForeignName == nil || course.Credit == nil ||
		*course.HoursTotal == nil || course.Assessment == nil || course.DepartmentName == nil ||
		course.LeaderName == nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Required fields cannot be empty.",
		})
		return
	}

	t, ok := (*course.HoursTotal).(int)
	if ok {
		var sum int
		if course.HoursLecture != nil {
			sum += *course.HoursLecture
		}
		if course.HoursPractices != nil {
			sum += *course.HoursPractices
		}
		if course.HoursExperiment != nil {
			sum += *course.HoursExperiment
		}
		if course.HoursComputer != nil {
			sum += *course.HoursComputer
		}
		if course.HoursSelf != nil {
			sum += *course.HoursSelf
		}
		if t != sum {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1001,
				Msg:  "Hours total is not equal to the sum of other hours.",
			})
			return
		}
	} else {
		if !(course.HoursLecture == nil && course.HoursPractices == nil && course.HoursExperiment == nil &&
			course.HoursComputer == nil && course.HoursSelf == nil) {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1001,
				Msg:  "Hours not set properly.",
			})
			return
		}
	}

	cs := &service.CourseService{}
	_, err := cs.CreateCourse(&course)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1002,
				Msg:  *course.Code + " already exists.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, Success(nil))
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	cs := &service.CourseService{}
	err = cs.DeleteCourse(uint(id))
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, Success(nil))
}

func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	var course model.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	cs := &service.CourseService{}
	err = cs.UpdateCourse(uint(id), &course)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1001,
				Msg:  "Course not found.",
			})
			return
		}
		if err.Error() == "errHoursTotal" {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1001,
				Msg:  "Hours not set properly.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(nil))
}

func (c *CourseController) GetCourseList(ctx *gin.Context) {
	var page model.Page
	page.Current, _ = strconv.Atoi(ctx.DefaultQuery("current", "1"))
	page.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("pageSize", "50"))
	if page.Current < 1 || page.PageSize < 1 {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Page meta error.",
		})
		return
	}

	var course model.Course
	var t string
	t = ctx.DefaultQuery("code", "")
	if t != "" {
		course.Code = new(string)
		*course.Code = t
	}
	t = ctx.DefaultQuery("name", "")
	if t != "" {
		course.Name = new(string)
		*course.Name = t
	}
	t = ctx.DefaultQuery("foreignName", "")
	if t != "" {
		course.ForeignName = new(string)
		*course.ForeignName = t
	}
	t = ctx.DefaultQuery("remark", "")
	if t != "" {
		course.Remark = new(string)
		*course.Remark = t
	}
	t = ctx.DefaultQuery("showRemark", "")
	if t != "" {
		course.ShowRemark = new(string)
		*course.ShowRemark = t
	}
	t = ctx.DefaultQuery("departmentName", "")
	if t != "" {
		course.DepartmentName = new(string)
		*course.DepartmentName = t
	}
	t = ctx.DefaultQuery("leaderName", "")
	if t != "" {
		course.LeaderName = new(string)
		*course.LeaderName = t
	}
	t = ctx.DefaultQuery("assessment", "")
	if t != "" {
		course.Assessment = new(string)
		*course.Assessment = t
	}
	credit, err := strconv.ParseFloat(ctx.DefaultQuery("credit", ""), 64)
	if err == nil {
		course.Credit = new(float64)
		*course.Credit = credit
	}

	cs := &service.CourseService{}
	var courses []model.Course
	err = cs.GetCourseList(&page, &course, &courses)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(gin.H{
		"page": page,
		"list": courses,
	}))
}

func (c *CourseController) UploadCourseFile(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *CourseController) GetCourseFileList(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}
