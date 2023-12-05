package controller

import (
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/config"
	"github.com/ljcbaby/plan/model"
	"github.com/ljcbaby/plan/service"
)

type Controller struct {
	Course  *CourseController
	Program *ProgramController
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Ljcbaby's plan backend.",
	})
}

func (c *Controller) DownloadFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	filepath := path.Join(config.Conf.Download.SavePath, filename)

	file, err := os.Open(filepath)
	if err != nil {
		ctx.String(http.StatusNotFound, "")
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	ctx.File(filepath)
}

func (c *Controller) GetTags(ctx *gin.Context) {
	var tags model.Tag

	ts := service.TagService{}
	if err := ts.GetTags(&tags); err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(tags))
}

func (c *Controller) PutTags(ctx *gin.Context) {
	var tags model.Tag
	if err := ctx.ShouldBindJSON(&tags); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	ts := service.TagService{}
	if err := ts.UpdateTags(&tags); err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(nil))
}

func Success(data interface{}) *model.Response {
	return &model.Response{
		Code: 0,
		Data: data,
		Msg:  "Success.",
	}
}

func returnMySQLError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, model.Response{
		Code: 1000,
		Msg:  "MySQL error.",
		Data: err.Error(),
	})
}
