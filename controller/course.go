package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/model"
)

type CourseController struct{}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
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
