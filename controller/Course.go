package controller

import (
	"io"
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
		course.DepartmentName == nil || course.LeaderName == nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Required fields cannot be empty.",
		})
		return
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
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(nil))
}

func (c *CourseController) GetCourseList(ctx *gin.Context) {
	cur := ctx.Query("current")
	if cur == "" {
		cur = "1"
	}
	pageSize := ctx.Query("pageSize")
	if pageSize == "" {
		pageSize = "50"
	}
	var page model.Page
	page.Current, _ = strconv.Atoi(cur)
	page.PageSize, _ = strconv.Atoi(pageSize)
	if page.Current < 1 || page.PageSize < 1 {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Page meta error.",
		})
		return
	}

	q := ctx.Query("q")
	if q != "" {
		cs := &service.CourseService{}
		var courses []model.Course

		err := cs.GetCourseListQuery(&page, &q, &courses)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, Success(gin.H{
			"page": page,
			"list": courses,
		}))
		return
	}

	var course model.Course
	var t string
	t = ctx.Query("code")
	if t != "" {
		course.Code = new(string)
		*course.Code = t
	}
	t = ctx.Query("name")
	if t != "" {
		course.Name = new(string)
		*course.Name = t
	}
	t = ctx.Query("foreignName")
	if t != "" {
		course.ForeignName = new(string)
		*course.ForeignName = t
	}
	t = ctx.Query("remark")
	if t != "" {
		course.Remark = new(string)
		*course.Remark = t
	}
	t = ctx.Query("showRemark")
	if t != "" {
		course.ShowRemark = new(string)
		*course.ShowRemark = t
	}
	t = ctx.Query("departmentName")
	if t != "" {
		course.DepartmentName = new(string)
		*course.DepartmentName = t
	}
	t = ctx.Query("leaderName")
	if t != "" {
		course.LeaderName = new(string)
		*course.LeaderName = t
	}
	t = ctx.Query("assessment")
	if t != "" {
		course.Assessment = new(string)
		*course.Assessment = t
	}
	credit, err := strconv.ParseFloat(ctx.Query("credit"), 64)
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
	file, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	var sum uint = 0
	var errs []string

	cs := &service.CourseService{}
	err = cs.ImportFile(file, &sum, &errs)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(gin.H{
		"success": sum,
		"err":     errs,
	}))
}

func (c *CourseController) GetCourseFileList(ctx *gin.Context) {
	flag := ctx.Query("content")

	cs := &service.CourseService{}
	var url string
	if flag == "all" {
		file, err := cs.ExportFile()
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
		url = "/v1/download/" + file
	} else {
		if err := cs.ReleaseTemplate(); err != nil {
			returnMySQLError(ctx, err)
			return
		}
		url = "/v1/download/course-template.xlsx"
	}
	ctx.JSON(http.StatusOK, Success(url))
}
