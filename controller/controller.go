package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/model"
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

func returnMySQLError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, model.Response{
		Code: 1000,
		Msg:  "MySQL error.",
		Data: err.Error(),
	})
}
