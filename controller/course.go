package controller

import (
	"net/http"
	"strconv"

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

	if course.Code == nil || course.Name == nil || course.ForeignName == nil || course.Credit == nil || course.HoursTotal == nil || course.Assessment == nil || course.DepartmentName == nil || course.LeaderName == nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Required fields cannot be empty.",
		})
		return
	}

	t, ok := (*course.HoursTotal).(int)
	if ok {
		if t != *course.HoursLecture+*course.HoursPractices+*course.HoursExperiment+*course.HoursComputer+*course.HoursSelf {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1001,
				Msg:  "Hours total is not equal to the sum of other hours.",
			})
			return
		}
	}

	cs := &service.CourseService{}
	_, err := cs.CreateCourse(&course)
	if err != nil {
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
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *CourseController) GetCourseList(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
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
